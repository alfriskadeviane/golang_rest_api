package buy

import (
	"time"

	"github.com/google/uuid"
)

type BuyResponse struct {
	ID     uuid.UUID `json:"id"`
	IdUser uint      `json:"id_user"`
	IdBuku uint      `json:"id_buku"`
	Jumlah int       `json:"jumlah"`
	Time   time.Time `json:"time"`
}
