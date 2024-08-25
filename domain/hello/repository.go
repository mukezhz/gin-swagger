package hello

type Greet struct {
	Message string `json:"message"`
	ID      string `json:"id"`
}

type Repository struct {
	greets []Greet
}

func NewRepository() *Repository {
	return &Repository{}
}

func (s *Repository) GetMessage() Model {
	return Model{Message: "Hello World"}
}

func (s *Repository) GetGreet(id string) Greet {
	for _, greet := range s.greets {
		if greet.ID == id {
			return greet
		}
	}
	return Greet{}
}

func (s *Repository) AddGreet(g Greet) {
	s.greets = append(s.greets, g)
}

func (s *Repository) UpdateGreet(id string, g Greet) {
	for i, greet := range s.greets {
		if greet.ID == id {
			s.greets[i] = g
		}
	}
}

func (s *Repository) DeleteGreet(id string) {
	for i, greet := range s.greets {
		if greet.ID == id {
			s.greets = append(s.greets[:i], s.greets[i+1:]...)
		}
	}
}
