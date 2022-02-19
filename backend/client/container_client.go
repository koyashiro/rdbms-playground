package client

import (
	"context"
	"errors"
	"io"
	"os"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"

	"github.com/koyashiro/rdbms-playground/backend/env"
)

type ContainerClient interface {
	GetAll(ctx context.Context) ([]types.Container, error)
	Get(ctx context.Context, id string) (*types.ContainerJSON, error)
	Create(ctx context.Context, name string, db string) (*types.ContainerJSON, error)
	Delete(ctx context.Context, id string) error
}

type containerClient struct {
	client     *client.Client
	sync.Mutex //TODO narrow the lock range
}

func NewContainerClient() ContainerClient {
	c, err := client.NewClientWithOpts(
		client.WithHost(client.DefaultDockerHost),
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		panic(err)
	}

	return &containerClient{client: c}
}

func (r *containerClient) GetAll(ctx context.Context) ([]types.Container, error) {
	r.Lock()
	defer r.Unlock()

	clo := types.ContainerListOptions{
		Filters: filters.NewArgs(filters.Arg("label", "type=playground")),
	}
	cl, err := r.client.ContainerList(ctx, clo)
	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (r *containerClient) Get(ctx context.Context, id string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	return r.get(ctx, id)
}

var limitsProcess = int64(50)

func (r *containerClient) Create(ctx context.Context, workspaceID string, db string) (*types.ContainerJSON, error) {
	r.Lock()
	defer r.Unlock()

	config, err := config(workspaceID, db)
	if err != nil {
		return nil, err
	}

	rhc := &container.HostConfig{
		AutoRemove: true,
		CapDrop:    []string{"fsetid", "kill", "setpcap", "net_raw", "sys_chroot", "mknod", "audit_write", "setfcap"},
		Resources: container.Resources{
			CPUQuota:  10000,
			CPUPeriod: 5000,
			Memory:    209_715_200,
			PidsLimit: &limitsProcess,
		},
	}

	reader, err := r.client.ImagePull(ctx, "docker.io/library/"+config.Image, types.ImagePullOptions{})
	if err != nil {
		return nil, err
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	ccb, err := r.client.ContainerCreate(ctx, config, rhc, nil, nil, workspaceID)
	if err != nil {
		return nil, err
	}

	if err := r.client.ContainerStart(ctx, ccb.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	err = r.client.NetworkConnect(ctx, env.Network, ccb.ID, nil)
	if err != nil {
		return nil, err
	}

	c, err := r.get(ctx, ccb.ID)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *containerClient) Delete(ctx context.Context, id string) error {
	r.Lock()
	defer r.Unlock()

	t := time.Second
	if err := r.client.ContainerStop(ctx, id, &t); err != nil {
		return err
	}

	return r.client.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         true,
	})
}

func (r *containerClient) get(ctx context.Context, id string) (*types.ContainerJSON, error) {
	c, err := r.client.ContainerInspect(ctx, id)
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
