package service

import (
	"github.com/google/uuid"

	"github.com/koyashiro/postgres-playground/backend/model"
	"github.com/koyashiro/postgres-playground/backend/repository"
)

type PlaygroundService interface {
	GetAll() ([]*model.Playground, error)
	Get(id string) (*model.Playground, error)
	Create(db string) (*model.Playground, error)
	Destroy(id string) error
	Execute(id string, query string) (string, error)
}

type PlaygroundServiceImpl struct {
	playgroundRepository repository.PlaygroundRepository
}

func NewPlaygroundService(repository *repository.PlaygroundRepository) PlaygroundService {
	return &PlaygroundServiceImpl{playgroundRepository: *repository}
}

func (s *PlaygroundServiceImpl) GetAll() ([]*model.Playground, error) {
	return s.playgroundRepository.GetAll()
}

func (s *PlaygroundServiceImpl) Get(id string) (*model.Playground, error) {
	return s.playgroundRepository.Get(id)
}

func (s *PlaygroundServiceImpl) Create(db string) (*model.Playground, error) {
	// TODO: create docker container

	p := &model.Playground{
		ID:      uuid.NewString(),
		DB:      db,
		Version: "13.0.0",
		Container: &model.Container{
			Hash:   "hash",
			Image:  "postgres",
			Status: "running",
			Port:   12345,
		},
	}

	if err := s.playgroundRepository.Set(p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *PlaygroundServiceImpl) Destroy(id string) error {
	// TODO: remove docker container

	return s.playgroundRepository.Delete(id)
}

func (s *PlaygroundServiceImpl) Execute(id string, query string) (string, error) {
	// TODO: execute query
	result := "XXXXXXXX"
	return result, nil
}
