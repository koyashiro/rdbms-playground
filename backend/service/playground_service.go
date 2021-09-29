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
	Execute(id string, query string) (*model.ExecuteResult, error)
}

type PlaygroundServiceImpl struct {
	cr repository.ContainerRepository
	rr repository.RDBMSRepository
}

func NewPlaygroundService(
	cr repository.ContainerRepository,
	rr repository.RDBMSRepository,
) PlaygroundService {
	return &PlaygroundServiceImpl{
		cr: cr,
		rr: rr,
	}
}

func (s *PlaygroundServiceImpl) GetAll() ([]*model.Playground, error) {
	containers, err := s.cr.GetAll()
	if err != nil {
		return nil, err
	}

	playgrounds := make([]*model.Playground, len(containers), len(containers))
	for i, container := range containers {
		c := model.NewContainerFromContainer(&container)
		playgrounds[i] = &model.Playground{
			ID:        c.Name,
			Container: c,
		}
	}

	return playgrounds, nil
}

func (s *PlaygroundServiceImpl) Get(id string) (*model.Playground, error) {
	cj, err := s.cr.Get(id)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	return &model.Playground{
		ID:        c.Name,
		Container: model.NewContainerFromContainerJSON(cj),
	}, nil
}

func (s *PlaygroundServiceImpl) Create(db string) (*model.Playground, error) {
	id := uuid.New().String()

	cj, err := s.cr.Create(id, db)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	p := &model.Playground{
		ID:        c.Name,
		Container: c,
	}

	return p, nil
}

func (s *PlaygroundServiceImpl) Destroy(id string) error {
	return s.cr.Delete(id)
}

func (s *PlaygroundServiceImpl) Execute(id string, query string) (*model.ExecuteResult, error) {
	cj, err := s.cr.Get(id)
	if err != nil {
		return nil, err
	}

	r, err := s.rr.Execute(cj, query)
	if err != nil {
		return nil, err
	}

	return r, nil
}
