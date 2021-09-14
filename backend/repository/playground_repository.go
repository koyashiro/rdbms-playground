package repository

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	"github.com/koyashiro/postgres-playground/backend/model"
)

type PlaygroundRepository interface {
	GetAll() ([]*model.Playground, error)
	Get(id string) (*model.Playground, error)
	Set(p *model.Playground) error
	Delete(id string) error
}

type PlaygroundRepositoryImpl struct {
	ctx    context.Context
	client *redis.Client
}

func NewPlaygroundRepository() PlaygroundRepository {
	ctx := context.Background()
	c := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	return &PlaygroundRepositoryImpl{
		ctx:    ctx,
		client: c,
	}
}

func (r *PlaygroundRepositoryImpl) GetAll() ([]*model.Playground, error) {
	ids, err := r.client.Keys(r.ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	ps := make([]*model.Playground, 0, len(ids))
	for _, id := range ids {
		p, err := r.Get(id)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}

	return ps, nil
}

func (r *PlaygroundRepositoryImpl) Get(id string) (*model.Playground, error) {
	b, err := r.client.Get(r.ctx, id).Bytes()
	if err != nil {
		return nil, err
	}

	var p *model.Playground
	if err = json.Unmarshal(b, &p); err != nil {
		return nil, err
	}

	return p, nil
}

func (r *PlaygroundRepositoryImpl) Set(p *model.Playground) error {
	j, err := json.Marshal(p)
	if err != nil {
		return err
	}

	return r.client.Set(r.ctx, p.ID, j, 0).Err()
}

func (r *PlaygroundRepositoryImpl) Delete(id string) error {
	return r.client.Del(r.ctx, id).Err()
}
