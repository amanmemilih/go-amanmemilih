package entities

type Document struct {
	Id         int
	DocumentC1 []string
	UserId     string
	Status     int
	CreatedAt  string
}

type DocumentVote struct {
	Id          *string
	DocumentId  *int
	CandidateId int
	TotalVote   int
}
