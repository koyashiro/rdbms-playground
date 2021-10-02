package service

import (
	"github.com/google/uuid"

	"github.com/koyashiro/rdbms-playground/backend/model"
	"github.com/koyashiro/rdbms-playground/backend/repository"
)

type WorkspaceService interface {
	GetAll() ([]*model.Workspace, error)
	Get(id string) (*model.Workspace, error)
	Create(db string) (*model.Workspace, error)
	Destroy(id string) error
	Execute(id string, query string) (*model.ExecuteResult, error)
}

type WorkspaceServiceImpl struct {
	containerRepository repository.ContainerRepository
	rdbmsRepository     repository.RDBMSRepository
}

func NewWorkspaceService(
	containerRepository repository.ContainerRepository,
	rdbmsRepository repository.RDBMSRepository,
) WorkspaceService {
	return &WorkspaceServiceImpl{
		containerRepository: containerRepository,
		rdbmsRepository:     rdbmsRepository,
	}
}

func (s *WorkspaceServiceImpl) GetAll() ([]*model.Workspace, error) {
	containers, err := s.containerRepository.GetAll()
	if err != nil {
		return nil, err
	}

	workspaces := make([]*model.Workspace, len(containers), len(containers))
	for i, container := range containers {
		c := model.NewContainerFromContainer(&container)
		workspaces[i] = &model.Workspace{
			ID:        c.Name,
			Container: c,
		}
	}

	return workspaces, nil
}

func (s *WorkspaceServiceImpl) Get(id string) (*model.Workspace, error) {
	cj, err := s.containerRepository.Get(id)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	return &model.Workspace{
		ID:        c.Name,
		Container: model.NewContainerFromContainerJSON(cj),
	}, nil
}

func (s *WorkspaceServiceImpl) Create(db string) (*model.Workspace, error) {
	id := uuid.New().String()

	cj, err := s.containerRepository.Create(id, db)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	p := &model.Workspace{
		ID:        c.Name,
		Container: c,
	}

	return p, nil
}

func (s *WorkspaceServiceImpl) Destroy(id string) error {
	return s.containerRepository.Delete(id)
}

func (s *WorkspaceServiceImpl) Execute(id string, query string) (*model.ExecuteResult, error) {
	cj, err := s.containerRepository.Get(id)
	if err != nil {
		return nil, err
	}

	r, err := s.rdbmsRepository.Execute(cj, query)
	if err != nil {
		return nil, err
	}

	return r, nil
}
