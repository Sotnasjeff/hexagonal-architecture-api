package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(persistence ProductPersistenceInterface) *ProductService {
	return &ProductService{
		Persistence: persistence,
	}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	prod, err := s.Persistence.Get(id)
	if err != nil {
		return nil, err
	}

	return prod, nil
}

func (s *ProductService) Create(name string, price float64) (ProductInterface, error) {
	prod := NewProduct()
	prod.Name = name
	prod.Price = price
	_, err := prod.IsValid()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(prod)
	if err != nil {
		return &Product{}, err
	}

	return result, nil
}

func (s *ProductService) Enable(product ProductInterface) (ProductInterface, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (s *ProductService) Disable(product ProductInterface) (ProductInterface, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.Persistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}
