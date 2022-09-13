package user

// "encoding/json"

type UsersRequest struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required" gorm:"uniq"`
	TempatLahir string `json:"tempatlahir" binding:"required"`
	TglLahir    string `json:"tanggal_lahir" binding:"required"`
	// SubTitle string `json:"sub_title"`
}
