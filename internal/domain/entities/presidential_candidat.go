package entities

type PresidentialCandidate struct {
	Id        int     `json:"id"`
	No        int     `json:"no"`
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
}
