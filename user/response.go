package user

// uuid "github.com/jackc/pgtype/ext/gofrs-uuid"

type UserResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	TempatLahir string `json:"tempatlahir"`
	TglLahir    string `json:"tanggal_lahir"`
}
