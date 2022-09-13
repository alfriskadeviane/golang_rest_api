package buy

import (
	// "mini-project/book"
	// "mini-project/user"
	"time"

	"github.com/google/uuid"
)

type Buy struct {

	// ID     string `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	// ID     uint `gorm:"primaryKey"`
	IdUser uint
	IdBuku uint
	Jumlah int
	Time   time.Time
	// User   user.User `gorm:"foreignKey:IdUser"`
	// Book   book.Book `gorm:"foreignKey:IdBook"`
}
