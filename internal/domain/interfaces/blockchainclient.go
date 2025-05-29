package interfaces

import "context"

type PresidentialVoteParams struct {
	CandidateId uint32 `json:"candidateId"`
	TotalVotes  uint32 `json:"totalVotes"`
}

type CreatePresidentialDocumentParams struct {
	CreatedAt     string                   `json:"createdAt"`
	DocumentC1Url []string                 `json:"documentC1Url"`
	UserId        uint32                   `json:"userId"`
	Vote          []PresidentialVoteParams `json:"vote"`
}

type BlockchainClient interface {
	CheckIfUserHasVoted(ctx context.Context, userId uint32) (bool, error)
	CreatePresidentialDocument(ctx context.Context, params CreatePresidentialDocumentParams) error
	DeleteAllData(ctx context.Context) error
	GetPresidentialDocuments(ctx context.Context) ([]interface{}, error)
	GetPresidentialVotes(ctx context.Context) ([]interface{}, error)
	GetTotalVotes(ctx context.Context) ([]interface{}, error)
}
