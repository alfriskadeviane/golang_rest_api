package buy

import (
	// "github.com/google/uuid"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]Buy, error)
	FindByID(ID uuid.UUID) (Buy, error)
	Create(buy Buy) (Buy, error)
	Update(buy Buy) (Buy, error)
	Delete(buy Buy) (Buy, error)
	GetPembelian() ([]Pembelian, error)
}
type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetPembelian() ([]Pembelian, error) {
	var pembelians []Pembelian
	err := r.db.Raw("select * from get_pembelian()").Scan(&pembelians).Error

	return pembelians, err
}

func (r *repository) FindAll() ([]Buy, error) {
	var buys []Buy
	err := r.db.Find(&buys).Error

	return buys, err
}

func (r *repository) FindByID(ID uuid.UUID) (Buy, error) {
	var buy Buy
	err := r.db.Find(&buy, ID).Error

	return buy, err
}

func (r *repository) Create(buy Buy) (Buy, error) {
	err := r.db.Create(&buy).Error

	return buy, err
}

func (r *repository) Update(buy Buy) (Buy, error) {
	err := r.db.Save(&buy).Error
	return buy, err
}

func (r *repository) Delete(buy Buy) (Buy, error) {
	err := r.db.Delete(&buy).Error
	return buy, err
}
