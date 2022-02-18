package service

import (
	"github.com/google/uuid"

	"github.com/koyashiro/rdbms-playground/backend/client"
	"github.com/koyashiro/rdbms-playground/backend/model"
)

type WorkspaceService interface {
	GetAll() ([]*model.Workspace, error)
	Get(id string) (*model.Workspace, error)
	Create(db string) (*model.Workspace, error)
	Delete(id string) error
	Execute(id string, query string) (*model.QueryResult, error)
}

type workspaceService struct {
	containerClient client.ContainerClient
	rdbmsClient     client.RDBMSClient
}

func NewWorkspaceService(
	containerClient client.ContainerClient,
	rdbmsClient client.RDBMSClient,
) WorkspaceService {
	return &workspaceService{
		containerClient: containerClient,
		rdbmsClient:     rdbmsClient,
	}
}

func (s *workspaceService) GetAll() ([]*model.Workspace, error) {
	containers, err := s.containerClient.GetAll()
	if err != nil {
		return nil, err
	}

	workspaces := make([]*model.Workspace, len(containers), len(containers))
	for i, container := range containers {
		c := model.NewContainerFromContainer(&container)
		workspaces[i] = &model.Workspace{
			ID: c.Name,
			DB: c.Image,
		}
	}

	return workspaces, nil
}

func (s *workspaceService) Get(id string) (*model.Workspace, error) {
	cj, err := s.containerClient.Get(id)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	return &model.Workspace{
		ID: c.Name,
		DB: c.Image,
	}, nil
}

func (s *workspaceService) Create(db string) (*model.Workspace, error) {
	id := uuid.New().String()

	cj, err := s.containerClient.Create(id, db)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	p := &model.Workspace{
		ID: c.Name,
		DB: c.Image,
	}

	return p, nil
}

func (s *workspaceService) Delete(id string) error {
	return s.containerClient.Delete(id)
}

func (s *workspaceService) Execute(id string, query string) (*model.QueryResult, error) {
	cj, err := s.containerClient.Get(id)
	if err != nil {
		return nil, err
	}

	r, err := s.rdbmsClient.Execute(cj, query)
	if err != nil {
		return nil, err
	}

	return r, nil
}
