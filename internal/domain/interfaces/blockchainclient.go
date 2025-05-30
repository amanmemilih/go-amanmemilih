package interfaces

import "context"

type PresidentialVoteParams struct {
	CandidateId uint32 `json:"candidateId"`
	TotalVotes  uint32 `json:"totalVotes"`
}

type DashboardResponse struct {
	NotUploaded uint8 `json:"not_uploaded"`
	Uploaded    uint8 `json:"uploaded"`
	Verified    uint8 `json:"verified"`
}

type CreatePresidentialDocumentParams struct {
	CreatedAt     string                   `json:"created_at"`
	DocumentC1Url []string                 `json:"document_c1_url"`
	UserId        uint32                   `json:"user_id"`
	Vote          []PresidentialVoteParams `json:"vote"`
}

type CheckDocumentResponse struct {
	ElectionType string `json:"election_type"`
	ID           uint32 `json:"id"`
	Name         string `json:"name"`
	Status       uint8  `json:"status"`
}

type VotePercentage struct {
	Name           string `json:"name"`
	No             uint8  `json:"no"`
	Image          string `json:"image"`
	ID             uint32 `json:"id"`
	VotePercentage string `json:"vote_percentage"`
}

type PresidentialDocumentDetailResponse struct {
	CreatedAt    string                           `json:"created_at"`
	DocumentC1   []string                         `json:"documents"`
	ElectionDate string                           `json:"election_date"`
	Status       uint8                            `json:"status"`
	Votes        []PresidentialDocumentDetailVote `json:"votes"`
	ElectionType string                           `json:"election_type"`
}

type PresidentialDocumentDetailVote struct {
	CandidateName string `json:"candidate_name"`
	CandidateNo   uint8  `json:"candidate_no"`
	TotalVotes    uint32 `json:"total_votes"`
}

type PresidentialDocument struct {
	CreatedAt  string   `json:"created_at"`
	DocumentC1 []string `json:"document_c1"`
	ID         uint32   `json:"id"`
	Status     uint8    `json:"status"`
	UserID     uint32   `json:"userId"`
}

type CreateDocumentParams struct {
	CreatedAt     string   `json:"created_at"`
	DocumentC1Url []string `json:"document_c1_url"`
	ElectionType  string   `json:"election_type"`
	UserID        uint32   `json:"user_id"`
}

type BlockchainClient interface {
	CheckDocument(ctx context.Context, userId uint32) ([]CheckDocumentResponse, error)
	CheckIfUserHasVoted(ctx context.Context, userId uint32) (bool, error)
	CreateDocument(ctx context.Context, params CreateDocumentParams) error
	CreatePresidentialDocument(ctx context.Context, params CreatePresidentialDocumentParams) error
	DeleteAllData(ctx context.Context) error
	GetDetailDocument(ctx context.Context, documentId uint32, electionType string) (*PresidentialDocumentDetailResponse, error)
	GetPresidentialDocuments(ctx context.Context) ([]PresidentialDocument, error)
	GetTotalVotes(ctx context.Context) ([]VotePercentage, error)
	VerifyDocument(ctx context.Context, documentId uint32, electionType string) (bool, error)
	GetDashboard(ctx context.Context, userId uint32) (*DashboardResponse, error)
}
