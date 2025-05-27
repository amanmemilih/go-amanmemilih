package entities

type Village struct {
	Id            int    `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	SubdistrictId int    `json:"subdistrict_id"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}
