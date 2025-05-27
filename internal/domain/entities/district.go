package entities

type District struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	ProvinceId int    `json:"province_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
