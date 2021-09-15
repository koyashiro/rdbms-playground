package model

import "database/sql"

type ExecuteResult struct {
	Columns []*Column       `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

// TODO: support Length, Precision and Scale
type Column struct {
	Name     string `json:"name"`
	Nullable bool   `json:"nullable"`
	// Length       int64  `json:"length"`
	DatabaseType string `json:"databaseType"`
	// Precision    int64  `json:"precision"`
	// Scale        int64  `json:"scale"`
}

func NewColumn(ct *sql.ColumnType) *Column {
	name := ct.Name()
	nullable, ok := ct.Nullable()
	if ok != false {
		panic("nullable is not supported")
	}
	databaseType := ct.DatabaseTypeName()

	return &Column{
		Name:         name,
		Nullable:     nullable,
		DatabaseType: databaseType,
	}
}
