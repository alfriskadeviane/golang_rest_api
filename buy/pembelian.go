package buy

import "time"

type Pembelian struct {
	Id_user       uint      `json:"Id_user"`
	Name          string    `json:"name"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Total_price   int       `json:"total_price"`
	Book_quantity int       `json:"book_quantity"`
	Payment_date  time.Time `json:"payment_date"`
}
