package entities

type PresidentialCandidatSummary struct {
	*PresidentialCandidate
	VotePercentage string `json:"vote_percentage"`
}
