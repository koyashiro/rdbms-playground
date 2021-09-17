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
	Name         string `json:"name"`
	DatabaseType string `json:"databaseType"`
	Nullable     bool   `json:"nullable,omitempty"`
	Length       int64  `json:"length,omitempty"`
	Precision    int64  `json:"precision,omitempty"`
	Scale        int64  `json:"scale,omitempty"`
}

func NewColumn(ct *sql.ColumnType) *ExportColumn {
	nullable, _ := ct.Nullable()
	length, _ := ct.Length()
	precision, scale, _ := ct.DecimalSize()

	return &ExportColumn{
		Name:         ct.Name(),
		DatabaseType: ct.DatabaseTypeName(),
		Nullable:     nullable,
		Length:       length,
		Precision:    precision,
		Scale:        scale,
	}
}
