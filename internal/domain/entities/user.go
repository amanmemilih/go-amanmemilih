package entities

type User struct {
	Id                 int     `json:"id"`
	Username           string  `json:"username"`
	UsernameVerifiedAt *string `json:"username_verified_at"`
	Password           *string `json:"password" `
	VillageId          int     `json:"village_id"`
	Address            string  `json:"address"`
	CreatedAt          *string `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
}
