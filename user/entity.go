package user

type User struct {
	// ID          string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	// ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ID          uint `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	TempatLahir string
	TglLahir    string
}
