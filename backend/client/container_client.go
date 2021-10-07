package client

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	"github.com/koyashiro/rdbms-playground/backend/env"
)

type ContainerClient interface {
	GetAll() ([]types.Container, error)
	Get(id string) (*types.ContainerJSON, error)
	Create(name string, db string) (*types.ContainerJSON, error)
	Delete(id string) error
}

type containerClient struct {
	ctx        context.Context
	client     *client.Client
	sync.Mutex //TODO narrow the lock range
}

func NewContainerClient() ContainerClient {
	ctx := context.Background()
	c, err := client.NewClientWithOpts(
		client.WithHost(client.DefaultDockerHost),
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		panic(err)
	}

	return &containerClient{ctx: ctx, client: c}
}

func (r *containerClient) GetAll() ([]types.Container, error) {
	r.Lock()
	defer r.Unlock()

	clo := types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "type=playground")),
	}
	cl, err := r.client.ContainerList(r.ctx, clo)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (r *containerClient) Get(id string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	return r.get(id)
}

func (r *containerClient) Create(workspaceID string, db string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	config, err := config(workspaceID, db)
	if err != nil {
		return nil, err
	}

	rhc := &container.HostConfig{
		CapDrop: []string{"fsetid", "kill", "setpcap", "net_raw", "sys_chroot", "mknod", "audit_write", "setfcap"},
	}
	ccb, err := r.client.ContainerCreate(r.ctx, config, rhc, nil, nil, workspaceID)
	if err != nil {
		return nil, err
	}

	if err := r.client.ContainerStart(r.ctx, ccb.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	err = r.client.NetworkConnect(r.ctx, env.Network, ccb.ID, nil)
	if err != nil {
		return nil, err
	}

	c, err := r.get(ccb.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *containerClient) Delete(id string) error {
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

func (r *containerClient) get(id string) (*types.ContainerJSON, error) {
	c, err := r.client.ContainerInspect(r.ctx, id)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func config(workspaceID string, db string) (*container.Config, error) {
	const password = "password"
	switch db {
	case "mysql":
		return &container.Config{
			Image: "mysql",
			Labels: map[string]string{
				"type": "playground",
				"wid":  workspaceID,
			},
			Env: []string{"MYSQL_ROOT_PASSWORD=" + password},
		}, nil
	case "mariadb":
		return &container.Config{
			Image: "mariadb",
			Labels: map[string]string{
				"type": "playground",
				"wid":  workspaceID,
			},
			Env: []string{"MARIADB_ROOT_PASSWORD=" + password},
		}, nil
	case "postgres":
		return &container.Config{
			Image: "postgres",
			Labels: map[string]string{
				"type": "playground",
				"wid":  workspaceID,
			},
			Env: []string{"POSTGRES_PASSWORD=" + password},
		}, nil
	default:
		return nil, errors.New("invalid db")
	}
}
