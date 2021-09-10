package repositories

import (
	"github.com/koyashiro/postgres-playground/backend/models"
)

// TODO: replace redis
var playgrounds map[string]*models.Playground = make(map[string]*models.Playground)

type PlaygroundRepository interface {
	GetAll() ([]*models.Playground, error)
	Get(id string) (*models.Playground, error)
	Append(playground *models.Playground) error
	Delete(id string) error
}

type PlaygroundRepositoryImpl struct{}

func NewPlaygroundRepository() PlaygroundRepository {
	return &PlaygroundRepositoryImpl{}
}

func (r *PlaygroundRepositoryImpl) GetAll() ([]*models.Playground, error) {
	s := make([]*models.Playground, len(playgrounds))
	for _, v := range playgrounds {
		s = append(s, v)
	}
	return s, nil
}

func (r *PlaygroundRepositoryImpl) Get(id string) (*models.Playground, error) {
	return playgrounds[id], nil
}

func (r *PlaygroundRepositoryImpl) Append(playground *models.Playground) error {
	playgrounds[playground.ID] = playground
	return nil
}

func (r *PlaygroundRepositoryImpl) Delete(id string) error {
	delete(playgrounds, id)
	return nil
}
