package hello

// Service handles the business logic of the module
type Service struct {
	// Add any dependencies here
	repo *Repository
}

// NewService creates a new instance of TestService
func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// GetMessage returns a greeting message
func (s *Service) GetMessage() Model {
	return s.repo.GetMessage()
}

// GetGreet returns a greet message by id
func (s *Service) GetGreet(id string) Greet {
	return s.repo.GetGreet(id)
}

// AddGreet adds a greet message
func (s *Service) AddGreet(g Greet) {
	s.repo.AddGreet(g)
}

// UpdateGreet updates a greet message by id
func (s *Service) UpdateGreet(id string, g Greet) {
	s.repo.UpdateGreet(id, g)
}

// DeleteGreet deletes a greet message by id
func (s *Service) DeleteGreet(id string) {
	s.repo.DeleteGreet(id)
}
