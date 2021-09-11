package repository

import (
	"context"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"

	"github.com/koyashiro/postgres-playground/backend/model"
)

type ContainerRepository interface {
	Get(id string) (*model.Container, error)
	Create() (*model.Container, error)
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

func (r *ContainerRepositoryImpl) Create() (*model.Container, error) {
	r.Lock()
	defer r.Unlock()

	ccb, err := r.create()
	if err != nil {
		return nil, err
	}

	err = r.start(ccb.ID)
	if err != nil {
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

	ports := make([]int, 0, len(c.NetworkSettings.Ports))
	for _, port := range c.NetworkSettings.Ports {
		if len(port) == 0 {
			continue
		}

		p, err := nat.ParsePort(port[0].HostPort)
		if err != nil {
			return nil, err
		}

		ports = append(ports, p)
	}

	return &model.Container{
		ID:     id,
		Image:  c.Image,
		Status: c.State.Status,
		Ports:  ports,
	}, nil
}

func (r *ContainerRepositoryImpl) create() (container.ContainerCreateCreatedBody, error) {
	cc := &container.Config{
		Image:  "postgres",
		Labels: map[string]string{"type": "playground"},
		Env:    []string{"POSTGRES_PASSWORD=password"},
	}

	hc := &container.HostConfig{PublishAllPorts: true}

	return r.client.ContainerCreate(r.ctx, cc, hc, nil, nil, "")
}

func (r *ContainerRepositoryImpl) start(id string) error {
	return r.client.ContainerStart(r.ctx, id, types.ContainerStartOptions{})
}
