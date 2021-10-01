package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/docker/docker/api/types"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/koyashiro/rdbms-playground/backend/model"
)

type RDBMSRepository interface {
	Execute(c *types.ContainerJSON, query string) (*model.ExecuteResult, error)
}

type RDBMSRepositoryImpl struct{}

func NewRDBMSRepository() RDBMSRepository {
	return &RDBMSRepositoryImpl{}
}

func (dr RDBMSRepositoryImpl) Execute(c *types.ContainerJSON, query string) (*model.ExecuteResult, error) {
	driverName, err := driverName(c.Config.Image)
	if err != nil {
		return nil, err
	}

	var host string
	if c.Name[0] == '/' {
		host = c.Name[1:]
	} else {
		host = c.Name
	}

	dataSourceName, err := dataSourceName(driverName, host)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rs, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rs.Close()

	columnTypes, err := rs.ColumnTypes()
	if err != nil {
		return nil, err
	}

	columns := make([]model.Column, len(columnTypes), len(columnTypes))
	for i := range columns {
		columns[i] = *model.NewColumn(columnTypes[i])
	}

	rows := make([][]interface{}, 0)

	for rs.Next() {
		row := make([]interface{}, len(columns))
		rowPtrs := make([]interface{}, len(columns))
		for i := range row {
			rowPtrs[i] = &row[i]
		}

		if err = rs.Scan(rowPtrs...); err != nil {
			return nil, err
		}

		for i := range row {
			switch r := row[i].(type) {
			case []byte:
				row[i] = string(r)
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, bool, float32, float64:
				row[i] = fmt.Sprint(r)
			}
		}

		rows = append(rows, row)
	}

	return &model.ExecuteResult{Columns: columns, Rows: rows}, nil
}

func driverName(image string) (string, error) {
	switch image {
	case "postgres":
		return "postgres", nil
	case "mysql", "mariadb":
		return "mysql", nil
	default:
		return "", errors.New("invalid image")
	}
}

func dataSourceName(driverName string, host string) (string, error) {
	switch driverName {
	case "postgres":
		return "host=" + host + " port=5432 user=postgres password=password dbname=postgres sslmode=disable", nil
	case "mysql", "mariadb":
		return "root:password@tcp(" + host + ":3306)/mysql", nil
	default:
		return "", errors.New("invalid driverName")
	}
}
