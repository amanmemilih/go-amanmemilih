package entities

type PresidentialDocument struct {
	Id         int    `json:"id"`
	DocumentC1 string `json:"document_c1"`
	UserId     int    `json:"user_id"`
	Status     int    `json:"status"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
