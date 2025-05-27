package entities

type PresidentialDocumentVote struct {
	Id                     int    `json:"id"`
	PresidentialDocumentId int    `json:"presidential_document_id"`
	PresidentialCandidatId int    `json:"presidential_candidat_id"`
	TotalVotes             int    `json:"total_votes"`
	CreatedAt              string `json:"created_at"`
	UpdatedAt              string `json:"updated_at"`
}
