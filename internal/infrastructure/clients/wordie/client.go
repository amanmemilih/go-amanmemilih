package wordie

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/zinct/amanmemilih/internal/domain/interfaces"
)

type Client struct {
	BaseURL string
	Timeout time.Duration
}

func NewClient() interfaces.WordClient {
	return &Client{
		BaseURL: "https://www.wordiebox.com/api/words",
		Timeout: 5 * time.Second,
	}
}

type wordResponse struct {
	Word     string `json:"Word"`
	Meaning  string `json:"Meaning"`
	Language string `json:"Language"`
}

func (c *Client) GetRandomWords(count int) ([]string, error) {
	client := http.Client{Timeout: c.Timeout}
	url := fmt.Sprintf("%s?country=indonesian&number=%d", c.BaseURL, count)

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to call wordie API: %w", err)
	}
	defer resp.Body.Close()

	var result []wordResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode wordie response: %w", err)
	}

	words := make([]string, len(result))
	for i, word := range result {
		words[i] = strings.ToLower(word.Word)
	}

	return words, nil
}
