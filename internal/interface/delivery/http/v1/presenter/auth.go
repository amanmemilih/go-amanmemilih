package presenter

type LoginResponse struct {
	User  UserLoginResponse `json:"user"`
	Token string            `json:"token"`
}

type UserLoginResponse struct {
	Id          int     `json:"id"`
	Username    string  `json:"username"`
	Address     string  `json:"address"`
	Village     string  `json:"village"`
	Province    string  `json:"province"`
	District    string  `json:"district"`
	Subdistrict string  `json:"subdistrict"`
	Region      string  `json:"region"`
	CreatedAt   *string `json:"created_at"`
}
