package subdomain

//Service is a default service
type Service struct{}

//NewService returns default Service
func NewService() *Service {
	return &Service{}
}

//List returns all entities
func (s *Service) List() []Subdomain {
	return allEntities
}

//Get returns entity by id
func (s *Service) Get(idx int) (*Subdomain, error) {
	return &allEntities[idx], nil
}
