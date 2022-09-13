package buy

import "time"

type BuysRequest struct {
	// ID     string    `json:"id" binding:"required"`

	IdUser uint `json:"id_user" binding:"required"`
	IdBuku uint `json:"id_buku" binding:"required"`
	Jumlah int  `json:"jumlah" binding:"required"`
	Time   time.Time
	// SubTitle string `json:"sub_title"`
}
