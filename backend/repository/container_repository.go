package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"

	"github.com/koyashiro/postgres-playground/backend/env"
	"github.com/koyashiro/postgres-playground/backend/model"
)

const password = "password"

var configs = map[string]*container.Config{
	"mysql": {
		Image:  "mysql",
		Labels: map[string]string{"type": "playground"},
		Env:    []string{"MYSQL_ROOT_PASSWORD=" + password},
	},
	"postgres": {
		Image:  "postgres",
		Labels: map[string]string{"type": "playground"},
		Env:    []string{"POSTGRES_PASSWORD=" + password},
	},
}

type ContainerRepository interface {
	Get(id string) (*model.Container, error)
	Create(name string, db string) (*model.Container, error)
	Delete(id string) error
}

type ContainerRepositoryImpl struct {
	ctx        context.Context
	client     *client.Client
	sync.Mutex //TODO narrow the lock range
}

func NewContainerRepository() (ContainerRepository, error) {
	ctx := context.Background()
	c, err := client.NewClientWithOpts(
		client.WithHost(client.DefaultDockerHost),
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		return nil, err
	}

	return &ContainerRepositoryImpl{ctx: ctx, client: c}, nil
}

func (r *ContainerRepositoryImpl) Get(id string) (*model.Container, error) {
	r.Lock()
	defer r.Unlock()

	return r.get(id)
}

func (r *ContainerRepositoryImpl) Create(name string, db string) (*model.Container, error) {
	r.Lock()
	defer r.Unlock()

	ccb, err := r.create(name, db)
	if err != nil {
		return nil, err
	}

	err = r.start(ccb.ID)
	if err != nil {
		return nil, err
	}

	if err := r.client.NetworkConnect(r.ctx, env.Network, ccb.ID, nil); err != nil {
		return nil, err
	}

	c, err := r.get(ccb.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *ContainerRepositoryImpl) Delete(id string) error {
	r.Lock()
	defer r.Unlock()

	t := time.Second
	if err := r.client.ContainerStop(r.ctx, id, &t); err != nil {
		return err
	}

	return r.client.ContainerRemove(r.ctx, id, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         true,
	})
}

func (r *ContainerRepositoryImpl) get(id string) (*model.Container, error) {
	c, err := r.client.ContainerInspect(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &model.Container{
		ID:     id,
		Image:  c.Image,
		Status: c.State.Status,
	}, nil
}

func (r *ContainerRepositoryImpl) create(name string, db string) (container.ContainerCreateCreatedBody, error) {
	c, ok := configs[db]
	if !ok {
		return container.ContainerCreateCreatedBody{}, errors.New("invalid db")
	}

	return r.client.ContainerCreate(r.ctx, c, nil, nil, nil, name)
}

func (r *ContainerRepositoryImpl) start(id string) error {
	return r.client.ContainerStart(r.ctx, id, types.ContainerStartOptions{})
}
