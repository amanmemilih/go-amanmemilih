package request

type LoginRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type RecoveryKeyRequest struct {
	Username string `json:"username" form:"username" binding:"required"`
}

type ChangePasswordRequest struct {
	Password string `json:"password" form:"password" binding:"required"`
	Phrase1  string `json:"phrase_1" form:"phrase_1" binding:"required"`
	Phrase2  string `json:"phrase_2" form:"phrase_2" binding:"required"`
	Phrase3  string `json:"phrase_3" form:"phrase_3" binding:"required"`
	Phrase4  string `json:"phrase_4" form:"phrase_4" binding:"required"`
	Phrase5  string `json:"phrase_5" form:"phrase_5" binding:"required"`
	Phrase6  string `json:"phrase_6" form:"phrase_6" binding:"required"`
	Phrase7  string `json:"phrase_7" form:"phrase_7" binding:"required"`
	Phrase8  string `json:"phrase_8" form:"phrase_8" binding:"required"`
	Phrase9  string `json:"phrase_9" form:"phrase_9" binding:"required"`
	Phrase10 string `json:"phrase_10" form:"phrase_10" binding:"required"`
	Phrase11 string `json:"phrase_11" form:"phrase_11" binding:"required"`
	Phrase12 string `json:"phrase_12" form:"phrase_12" binding:"required"`
}
