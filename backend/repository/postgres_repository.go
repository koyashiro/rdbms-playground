package repository

import (
	_ "github.com/lib/pq"

	"github.com/koyashiro/postgres-playground/backend/model"
)

type PostgresRepositoryImpl struct{}

func NewPostgresRepository() DBRepository {
	return &PostgresRepositoryImpl{}
}

func (r *PostgresRepositoryImpl) Execute(name string, query string) (*model.ExecuteResult, error) {
	driverName := "postgres"
	dataSourceName := "host=" + name + " port=5432 user=postgres password=password dbname=postgres sslmode=disable"

	return Execute(driverName, dataSourceName, query)
}
