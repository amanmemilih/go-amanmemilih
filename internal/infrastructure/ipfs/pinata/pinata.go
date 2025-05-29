package pinata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/zinct/amanmemilih/config"
	"github.com/zinct/amanmemilih/internal/infrastructure/ipfs"
	"github.com/zinct/amanmemilih/pkg/logger"
)

type Pinata struct {
	apiKey    string
	apiSecret string
	baseURL   string
	client    *http.Client
	logger    *logger.Logger
}

type pinataResponse struct {
	IpfsHash string `json:"IpfsHash"`
}

type pinataMetadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Group       string            `json:"group"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

func NewPinata(cfg *config.Config, logger *logger.Logger) ipfs.IPFS {
	return &Pinata{
		apiKey:    cfg.Pinata.APIKey,
		apiSecret: cfg.Pinata.APISecret,
		baseURL:   "https://api.pinata.cloud",
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		logger: logger,
	}
}

func (p *Pinata) Upload(ctx context.Context, content string, metadata *ipfs.Metadata) (string, error) {
	// Create a new multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Use metadata name as filename
	fileName := fmt.Sprintf("documentc1-%s", uuid.New().String())

	// Get content type from file extension
	contentType := "text/plain"
	ext := strings.ToLower(filepath.Ext(fileName))
	switch ext {
	case ".jpg", ".jpeg":
		contentType = "image/jpeg"
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".pdf":
		contentType = "application/pdf"
	}

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", fmt.Errorf("failed to create form file: %w", err)
	}

	if _, err := io.WriteString(part, content); err != nil {
		return "", fmt.Errorf("failed to write content: %w", err)
	}

	// Add metadata if provided
	if metadata != nil {
		pinataMetadata := pinataMetadata{
			Name:        fileName, // Use the same filename in metadata
			Description: metadata.Description,
			Group:       metadata.Group,
			Attributes: map[string]string{
				"contentType": contentType,
			},
		}

		// Add any additional attributes
		for k, v := range metadata.Attributes {
			pinataMetadata.Attributes[k] = v
		}

		metadataBytes, err := json.Marshal(pinataMetadata)
		if err != nil {
			return "", fmt.Errorf("failed to marshal metadata: %w", err)
		}

		if err := writer.WriteField("pinataMetadata", string(metadataBytes)); err != nil {
			return "", fmt.Errorf("failed to write metadata: %w", err)
		}
	}

	// Close the writer
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	// Create the request
	req, err := http.NewRequestWithContext(ctx, "POST", p.baseURL+"/pinning/pinFileToIPFS", body)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("pinata_api_key", p.apiKey)
	req.Header.Set("pinata_secret_api_key", p.apiSecret)

	// Send the request
	resp, err := p.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("pinata api error: %s", string(respBody))
	}

	// Parse the response
	var pinataResp pinataResponse
	if err := json.Unmarshal(respBody, &pinataResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	return pinataResp.IpfsHash, nil
}

func (p *Pinata) UploadMultiple(ctx context.Context, contents []string, group string) ([]string, error) {
	hashes := make([]string, len(contents))
	errChan := make(chan error, len(contents))
	var wg sync.WaitGroup

	for i, content := range contents {
		wg.Add(1)
		go func(index int, content string) {
			defer wg.Done()

			// Generate random UUID for filename
			randomUUID := uuid.New().String()

			// Create metadata for each file
			metadata := &ipfs.Metadata{
				Name:        fmt.Sprintf("documentc1-%s", randomUUID),
				Description: fmt.Sprintf("Document %d in group %s", index+1, group),
				Group:       group,
				Attributes: map[string]string{
					"index": fmt.Sprintf("%d", index+1),
					"group": group,
				},
			}

			hash, err := p.Upload(ctx, content, metadata)
			if err != nil {
				errChan <- fmt.Errorf("failed to upload document %d: %w", index, err)
				return
			}
			hashes[index] = hash
		}(i, content)
	}

	// Wait for all uploads to complete
	wg.Wait()
	close(errChan)

	// Check for any errors
	if len(errChan) > 0 {
		var errMsgs []string
		for err := range errChan {
			errMsgs = append(errMsgs, err.Error())
		}
		return nil, fmt.Errorf("upload errors: %s", strings.Join(errMsgs, "; "))
	}

	return hashes, nil
}
