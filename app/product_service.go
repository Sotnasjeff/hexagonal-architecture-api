package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	prod, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return prod, nil
}
