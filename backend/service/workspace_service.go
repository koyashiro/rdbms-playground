package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/koyashiro/rdbms-playground/backend/client"
	"github.com/koyashiro/rdbms-playground/backend/model"
)

type WorkspaceService interface {
	GetAll(ctx context.Context) ([]*model.Workspace, error)
	Get(ctx context.Context, id string) (*model.Workspace, error)
	Create(ctx context.Context, db string) (*model.Workspace, error)
	Delete(ctx context.Context, id string) error
	Execute(ctx context.Context, id string, query string) (*model.QueryResult, error)
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

func (s *workspaceService) GetAll(ctx context.Context) ([]*model.Workspace, error) {
	containers, err := s.containerClient.GetAll(ctx)
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

func (s *workspaceService) Get(ctx context.Context, id string) (*model.Workspace, error) {
	cj, err := s.containerClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	c := model.NewContainerFromContainerJSON(cj)
	return &model.Workspace{
		ID: c.Name,
		DB: c.Image,
	}, nil
}

func (s *workspaceService) Create(ctx context.Context, db string) (*model.Workspace, error) {
	id := uuid.New().String()

	cj, err := s.containerClient.Create(ctx, id, db)
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

func (s *workspaceService) Delete(ctx context.Context, id string) error {
	return s.containerClient.Delete(ctx, id)
}

func (s *workspaceService) Execute(ctx context.Context, id string, query string) (*model.QueryResult, error) {
	cj, err := s.containerClient.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	r, err := s.rdbmsClient.Execute(cj, query)
	if err != nil {
		return nil, err
	}

	return r, nil
}
