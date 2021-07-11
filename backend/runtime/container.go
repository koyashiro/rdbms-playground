package runtime

import (
	"context"
	"errors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/docker/go-connections/nat"

	"github.com/docker/docker/api/types/mount"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/client"
)

var ErrDBNotFound = errors.New("database is not found")

var Logger *log.Logger

func logf(format string, v ...interface{}) {
	if Logger == nil {
		log.Printf(format, v...)
		return
	}
	Logger.Printf(format, v...)
}

type DBManager interface {
	Create(id string) (*DB, error)
	Destroy(id string) error
	Stat(id string) (*DB, error)
}

type DBManage struct {
	usePackage string
	resolve    map[string]*DB
	client     *client.Client
	sync.Mutex //TODO narrow the lock range
}

func NewDBManage(host string) DBManager {
	cli, err := client.NewClientWithOpts(
		client.WithHost(host),
		client.WithAPIVersionNegotiation(),
	)

	if err != nil {
		logf("%s", err)
		return nil
	}

	return &DBManage{
		usePackage: "postgres",
		client:     cli,
		resolve:    map[string]*DB{},
	}
}

type DB struct {
	ID     string
	Hash   string
	Status string
	Port   int
}

func (d *DBManage) Create(id string) (*DB, error) {
	d.Lock()
	defer d.Unlock()

	ctx := context.Background()

	_, err := d.client.ImagePull(ctx, d.usePackage, types.ImagePullOptions{})
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	cc := &container.Config{
		Image:  d.usePackage,
		Labels: map[string]string{"type": "playground"},
		Env:    []string{"POSTGRES_PASSWORD=password"},
	}

	dir, err := os.Getwd()
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	hc := &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: dir + "/runtime/misc/init_db.sh",
				Target: "/docker-entrypoint-initdb.d/init_db.sh",
			},
		},
		PublishAllPorts: true,
	}

	ccb, err := d.client.ContainerCreate(
		ctx,
		cc,
		hc,
		nil,
		nil,
		"",
	)
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	err = d.client.ContainerStart(ctx, ccb.ID, types.ContainerStartOptions{})
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	//TODO implements blocking container running

	stat, err := d.client.ContainerInspect(ctx, ccb.ID)
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	db := &DB{
		ID:     id,
		Hash:   ccb.ID,
		Status: stat.State.Status,
	}

	for _, port := range stat.NetworkSettings.Ports {
		if len(port) > 0 {
			// when enable ipv6 return multi ports
			i, err := nat.ParsePort(port[0].HostPort)
			if err != nil {
				logf("%s", err)
				return nil, err
			}
			db.Port = i
		}
		break
	}

	d.resolve[id] = db

	return db, nil
}

func (d *DBManage) Destroy(id string) error {

	db, err := d.resolveDB(id)
	if err != nil {
		logf("%s", err)
		return err
	}

	d.Lock()
	defer d.Unlock()

	ctx := context.Background()
	t := time.Second

	err = d.client.ContainerStop(ctx, db.Hash, &t)
	if err != nil {
		logf("%s", err)
		return err
	}

	err = d.client.ContainerRemove(ctx, db.Hash, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		RemoveLinks:   false,
		Force:         true,
	})

	if err != nil {
		logf("%s", err)
		return err
	}

	delete(d.resolve, id)

	return nil
}

func (d *DBManage) Stat(id string) (*DB, error) {
	d.Lock()
	defer d.Unlock()

	db, err := d.resolveDB(id)
	if err != nil {
		logf("%s", err)
		return nil, err
	}
	ctx := context.Background()

	stat, err := d.client.ContainerInspect(ctx, id)
	if err != nil {
		logf("%s", err)
		return nil, err
	}

	db.Status = stat.State.Status

	return db, nil
}

func (d *DBManage) resolveDB(id string) (*DB, error) {
	d.Lock()
	defer d.Unlock()
	db, ok := d.resolve[id]
	if !ok {
		logf("db not found got id %s", id)
		return nil, ErrDBNotFound
	}

	return db, nil
}
