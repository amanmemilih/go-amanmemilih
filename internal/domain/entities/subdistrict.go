package entities

type Subdistrict struct {
	Id         int    `json:"id"`
	Code       string `json:"code"`
	Name       string `json:"name"`
	DistrictId int    `json:"district_id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
