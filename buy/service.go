package buy

import "github.com/google/uuid"

type Service interface {
	FindAll() ([]Buy, error)
	FindByID(ID uuid.UUID) (Buy, error)
	Create(buy BuysRequest) (Buy, error)
	Update(ID uuid.UUID, buy BuysRequest) (Buy, error)
	Delete(ID uuid.UUID) (Buy, error)
	GetPembelian() ([]Pembelian, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetPembelian() ([]Pembelian, error) {
	pembelians, err := s.repository.GetPembelian()
	return pembelians, err
}

func (s *service) FindAll() ([]Buy, error) {
	buys, err := s.repository.FindAll()
	return buys, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID uuid.UUID) (Buy, error) {
	buy, err := s.repository.FindByID(ID)
	return buy, err
}

func (s *service) Create(buyrequest BuysRequest) (Buy, error) {
	buy := Buy{

		IdUser: buyrequest.IdUser,
		IdBuku: buyrequest.IdBuku,
		Jumlah: buyrequest.Jumlah,
		Time:   buyrequest.Time,
	}
	newBuy, err := s.repository.Create(buy)
	return newBuy, err
}

func (s *service) Update(ID uuid.UUID, buyrequest BuysRequest) (Buy, error) {
	buy, err := s.repository.FindByID(ID)

	buy.IdUser = buyrequest.IdUser
	buy.IdBuku = buyrequest.IdBuku
	buy.Jumlah = buyrequest.Jumlah
	buy.Time = buyrequest.Time

	newBuy, err := s.repository.Update(buy)
	return newBuy, err
}

func (s *service) Delete(ID uuid.UUID) (Buy, error) {
	buy, err := s.repository.FindByID(ID)

	newBuy, err := s.repository.Delete(buy)
	return newBuy, err
}
