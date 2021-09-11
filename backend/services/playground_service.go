package services

import (
	"github.com/google/uuid"

	"github.com/koyashiro/postgres-playground/backend/models"
	"github.com/koyashiro/postgres-playground/backend/repositories"
)

type PlaygroundService interface {
	GetAll() ([]*models.Playground, error)
	Get(id string) (*models.Playground, error)
	Create(db string) (*models.Playground, error)
	Destroy(id string) error
	Execute(id string, query string) (string, error)
}

type PlaygroundServiceImpl struct {
	playgroundRepository repositories.PlaygroundRepository
}

func NewPlaygroundService(repository *repositories.PlaygroundRepository) PlaygroundService {
	return &PlaygroundServiceImpl{playgroundRepository: *repository}
}

func (s *PlaygroundServiceImpl) GetAll() ([]*models.Playground, error) {
	return s.playgroundRepository.GetAll()
}

func (s *PlaygroundServiceImpl) Get(id string) (*models.Playground, error) {
	return s.playgroundRepository.Get(id)
}

func (s *PlaygroundServiceImpl) Create(db string) (*models.Playground, error) {
	// TODO: create docker container

	p := &models.Playground{
		ID:      uuid.NewString(),
		DB:      db,
		Version: "13.0.0",
		Container: &models.Container{
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
