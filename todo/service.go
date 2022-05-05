package todo

type Service interface {
	CreateTodo(input InputTodo) (Todo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTodo(input InputTodo) (Todo, error) {
	todo := Todo{}
	todo.Name = input.Name
	todo.Description = input.Description

	newTodo, err := s.repository.Save(todo)
	if err != nil {
		return newTodo, err
	}

	return newTodo, nil
}

