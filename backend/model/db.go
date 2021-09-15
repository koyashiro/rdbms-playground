package model

import (
	"bytes"
	"database/sql"
	"fmt"
	"strconv"
)

type ExecuteResult struct {
	Columns []*Column       `json:"columns"`
	Rows    [][]interface{} `json:"rows"`
}

// sql.ColumnType wrapper
type Column struct {
	*sql.ColumnType
}

func NewColumn(ct *sql.ColumnType) *Column {
	return &Column{ColumnType: ct}
}

func (c *Column) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer

	if err := b.WriteByte('{'); err != nil {
		return nil, err
	}

	if _, err := b.WriteString("\"name\":"); err != nil {
		return nil, err
	}

	if _, err := b.WriteString("\"" + c.Name() + "\""); err != nil {
		return nil, err
	}

	if err := b.WriteByte(','); err != nil {
		return nil, err
	}

	if _, err := b.WriteString("\"databaseType\":"); err != nil {
		return nil, err
	}

	if _, err := b.WriteString("\"" + c.DatabaseTypeName() + "\""); err != nil {
		return nil, err
	}

	if nullable, ok := c.Nullable(); ok {
		if err := b.WriteByte(','); err != nil {
			return nil, err
		}

		if _, err := b.WriteString("\"nullable\":"); err != nil {
			return nil, err
		}

		if _, err := b.WriteString(strconv.FormatBool(nullable)); err != nil {
			return nil, err
		}
	}

	if length, ok := c.Length(); ok {
		if err := b.WriteByte(','); err != nil {
			return nil, err
		}

		if _, err := b.WriteString("\"length\":"); err != nil {
			return nil, err
		}

		if _, err := b.WriteString(fmt.Sprint(length)); err != nil {
			return nil, err
		}
	}

	if precision, scale, ok := c.DecimalSize(); ok {
		if err := b.WriteByte(','); err != nil {
			return nil, err
		}

		if _, err := b.WriteString("\"precision\":"); err != nil {
			return nil, err
		}

		if _, err := b.WriteString(fmt.Sprint(precision)); err != nil {
			return nil, err
		}

		if err := b.WriteByte(','); err != nil {
			return nil, err
		}

		if _, err := b.WriteString("\"scale\":"); err != nil {
			return nil, err
		}

		if _, err := b.WriteString(fmt.Sprint(scale)); err != nil {
			return nil, err
		}
	}

	if err := b.WriteByte('}'); err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}
