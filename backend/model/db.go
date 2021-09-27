package model

import (
	"database/sql"
)

type ExecuteResult struct {
	Columns []*Column       `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

type Column struct {
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
	Nullable     *bool  `json:"nullable,omitempty"`
	Length       *int64 `json:"length,omitempty"`
	Precision    *int64 `json:"precision,omitempty"`
	Scale        *int64 `json:"scale,omitempty"`
}

func NewExportColumn(ct *sql.ColumnType) *Column {
	var nullable *bool
	if n, ok := ct.Nullable(); ok {
		nullable = &n
	}

	var length *int64
	if l, ok := ct.Length(); ok {
		length = &l
	}

	var precision, scale *int64
	if p, s, ok := ct.DecimalSize(); ok {
		precision, scale = &p, &s
	}

	databaseType := ct.DatabaseTypeName()

	return &Column{
		Name:         ct.Name(),
		DatabaseType: databaseType,
		Nullable:     nullable,
		Length:       length,
		Precision:    precision,
		Scale:        scale,
	}
}
