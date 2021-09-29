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
	cr repository.ContainerRepository
	rr repository.RDBMSRepository
}

func NewWorkspaceService(
	cr repository.ContainerRepository,
	rr repository.RDBMSRepository,
) WorkspaceService {
	return &WorkspaceServiceImpl{
		cr: cr,
		rr: rr,
	}
}

func (s *WorkspaceServiceImpl) GetAll() ([]*model.Workspace, error) {
	containers, err := s.cr.GetAll()
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
	cj, err := s.cr.Get(id)
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

	cj, err := s.cr.Create(id, db)
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
	return s.cr.Delete(id)
}

func (s *WorkspaceServiceImpl) Execute(id string, query string) (*model.ExecuteResult, error) {
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
