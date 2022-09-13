package user

type Service interface {
	FindAll() ([]User, error)
	FindByID(ID int) (User, error)
	Create(user UsersRequest) (User, error)
	Update(ID int, user UsersRequest) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
	// return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)
	return user, err
}

func (s *service) Create(userRequest UsersRequest) (User, error) {
	// price, _ := userRequest.Price.Int64()
	// rating, _ := userRequest.Rating.Int64()

	user := User{
		Name:        userRequest.Name,
		Email:       userRequest.Email,
		TempatLahir: userRequest.TempatLahir,
		TglLahir:    userRequest.TglLahir,
	}
	newuser, err := s.repository.Create(user)
	return newuser, err
}

func (s *service) Update(ID int, userRequest UsersRequest) (User, error) {
	user, err := s.repository.FindByID(ID)

	// price, _ := userRequest.Price.Int64()
	// rating, _ := userRequest.Rating.Int64()

	user.Name = userRequest.Name
	user.Email = userRequest.Email
	user.TempatLahir = userRequest.TempatLahir
	user.TglLahir = userRequest.TglLahir

	newuser, err := s.repository.Update(user)
	return newuser, err
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	newuser, err := s.repository.Delete(user)
	return newuser, err
}
