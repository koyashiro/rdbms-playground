package model

import (
	"database/sql"
)

type ExecuteResult struct {
	Columns []*ExportColumn `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

// sql.ColumnType wrapper
type Column struct {
	*sql.ColumnType
}

type ExportColumn struct {
	Name         string      `json:"name"`
	DatabaseType interface{} `json:"databaseType"`
	Nullable     interface{} `json:"nullable,omitempty"`
	Length       interface{} `json:"length,omitempty"`
	Precision    interface{} `json:"precision,omitempty"`
	Scale        interface{} `json:"scale,omitempty"`
}

func NewColumn(ct *sql.ColumnType) *ExportColumn {
	var nullable interface{}
	if n, ok := ct.Nullable(); ok {
		nullable = n
	} else {
		nullable = interface{}(nil)
	}

	var length interface{}
	if l, ok := ct.Length(); ok {
		length = l
	} else {
		length = interface{}(nil)
	}

	var precision, scale interface{}
	if p, s, ok := ct.DecimalSize(); ok {
		precision, scale = p, s
	} else {
		precision, scale = interface{}(nil), interface{}(nil)
	}

	return &ExportColumn{
		Name:         ct.Name(),
		DatabaseType: ct.DatabaseTypeName(),
		Nullable:     nullable,
		Length:       length,
		Precision:    precision,
		Scale:        scale,
	}
}
