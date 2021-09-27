package repository

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/koyashiro/postgres-playground/backend/model"
)

type MysqlRepositoryImpl struct{}

func NewMysqlRepository() DBRepository {
	return &MysqlRepositoryImpl{}
}

func (r *MysqlRepositoryImpl) Execute(name string, query string) (*model.ExecuteResult, error) {
	driverName := "mysql"
	dataSourceName := "host=" + name + " port=3306 user=root password=password dbname=mysql sslmode=disable"

	return Execute(driverName, dataSourceName, query)
}
