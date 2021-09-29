package repository

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	"github.com/koyashiro/postgres-playground/backend/env"
)

type ContainerRepository interface {
	GetAll() ([]types.Container, error)
	Get(id string) (*types.ContainerJSON, error)
	Create(name string, db string) (*types.ContainerJSON, error)
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

func (r *ContainerRepositoryImpl) GetAll() ([]types.Container, error) {
	r.Lock()
	defer r.Unlock()

	return r.getAll()
}

func (r *ContainerRepositoryImpl) Get(id string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	return r.get(id)
}

func (r *ContainerRepositoryImpl) Create(playgroundID string, db string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	ccb, err := r.create(playgroundID, db)
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

func (r *ContainerRepositoryImpl) getAll() ([]types.Container, error) {
	clo := types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "type=playground")),
	}
	cl, err := r.client.ContainerList(r.ctx, clo)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (r *ContainerRepositoryImpl) get(id string) (*types.ContainerJSON, error) {
	c, err := r.client.ContainerInspect(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

var restrictHostConfig = &container.HostConfig{
	CapDrop: []string{"fsetid", "kill", "setpcap", "net_raw", "sys_chroot", "mknod", "audit_write", "setfcap"},
}

func config(playgroundID string, db string) (*container.Config, error) {
	const password = "password"
	switch db {
	case "mysql":
		return &container.Config{
			Image: "mysql",
			Labels: map[string]string{
				"type": "playground",
				"pgid": playgroundID,
			},
			Env: []string{"MYSQL_ROOT_PASSWORD=" + password},
		}, nil
	case "postgres":
		return &container.Config{
			Image: "postgres",
			Labels: map[string]string{
				"type": "playground",
				"pgid": playgroundID,
			},
			Env: []string{"POSTGRES_PASSWORD=" + password},
		}, nil
	default:
		return nil, errors.New("invalid db")
	}
}

func (r *ContainerRepositoryImpl) create(playgroundID string, db string) (container.ContainerCreateCreatedBody, error) {
	c, err := config(playgroundID, db)
	if err != nil {
		return container.ContainerCreateCreatedBody{}, err
	}

	return r.client.ContainerCreate(r.ctx, c, restrictHostConfig, nil, nil, playgroundID)
}

func (r *ContainerRepositoryImpl) start(id string) error {
	return r.client.ContainerStart(r.ctx, id, types.ContainerStartOptions{})
}
