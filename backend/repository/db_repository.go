package repository

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/koyashiro/postgres-playground/backend/model"
)

func dataSourceName(driverName string, host string) (string, error) {
	switch driverName {
	case "postgres":
		return "host=" + host + " port=5432 user=postgres password=password dbname=postgres sslmode=disable", nil
	case "mysql":
		return "root:password@tcp(" + host + ":3306)/mysql", nil
	default:
		return "", errors.New("invalid driverName")
	}
}

type DBRepository interface {
	Execute(playground *model.Playground, query string) (*model.ExecuteResult, error)
}

type DBRepositoryImpl struct{}

func NewDBRepository() DBRepository {
	return &DBRepositoryImpl{}
}

func (dr DBRepositoryImpl) Execute(playground *model.Playground, query string) (*model.ExecuteResult, error) {
	driverName := playground.DB
	dataSourceName, err := dataSourceName(driverName, playground.ID)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	types, err := rows.ColumnTypes()
	if err != nil {
		return nil, err
	}

	values := make([][]interface{}, 0)

	ptrs := make([]interface{}, len(types))

	for rows.Next() {
		row := make([]interface{}, len(types))
		for i := range row {
			ptrs[i] = &row[i]
		}

		if err = rows.Scan(ptrs...); err != nil {
			return nil, err
		}

		values = append(values, row)
	}

	columns := make([]*model.Column, len(types), len(types))
	for i := range columns {
		columns[i] = model.NewExportColumn(types[i])
	}

	return &model.ExecuteResult{Columns: columns, Rows: values}, nil
}
