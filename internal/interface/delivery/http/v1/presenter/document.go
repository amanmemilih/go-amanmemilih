package presenter

type DocumentUserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	VillageId int    `json:"village_id"`
}
