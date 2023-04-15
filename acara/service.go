package acara

type Service interface {
	GetAllAcaras() ([]Acara, error)
	GetAcaraByID(input GetAcaraDetailInput) (Acara, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllAcaras() ([]Acara, error) {
	acaras, err := s.repository.FindAll()
	if err != nil {
		return acaras, err
	}

	return acaras, nil
}

func (s *service) GetAcaraByID(input GetAcaraDetailInput) (Acara, error) {
	acara, err := s.repository.FindByID(input.ID)

	if err != nil {
		return acara, err
	}

	return acara, nil
}
