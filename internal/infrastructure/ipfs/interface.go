package ipfs

import "context"

type Metadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Group       string            `json:"group"`
	Attributes  map[string]string `json:"attributes,omitempty"`
}

type IPFS interface {
	Upload(ctx context.Context, content string, metadata *Metadata) (string, error)
	UploadMultiple(ctx context.Context, contents []string, group string) ([]string, error)
}
