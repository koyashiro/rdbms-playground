package repository

import (
	"database/sql"

	"github.com/koyashiro/postgres-playground/backend/model"
)

type DBRepository interface {
	Execute(name string, query string) (*model.ExecuteResult, error)
}

func Execute(driverName string, dataSourceName string, query string) (*model.ExecuteResult, error) {
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
