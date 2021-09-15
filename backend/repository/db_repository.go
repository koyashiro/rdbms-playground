package repository

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/lib/pq"
)

type DBRepository interface {
	Execute(namw string, query string) (string, error)
}

type PostgresRepositoryImpl struct{}

func NewPostgresRepository() DBRepository {
	return &PostgresRepositoryImpl{}
}

func (r *PostgresRepositoryImpl) Execute(name string, query string) (string, error) {
	db, err := sql.Open("postgres", "host="+name+" port=5432 user=postgres password=password dbname=postgres sslmode=disable")
	if err != nil {
		return "", err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	types, err := rows.ColumnTypes()
	if err != nil {
		return "", err
	}

	valuesPtrs := make([]interface{}, len(types))

	var sb strings.Builder
	for rows.Next() {
		values := make([]interface{}, len(types))
		for i := range values {
			valuesPtrs[i] = &values[i]
		}

		if err = rows.Scan(valuesPtrs...); err != nil {
			return "", err
		}

		types[0].DatabaseTypeName()

		valueStrs := make([]string, len(types), len(types))
		for i, value := range values {
			switch v := value.(type) {
			case int32, int64, float32, float64:
				valueStrs[i] = fmt.Sprint(v)
			case []byte:
				valueStrs[i] = string(v)
			case string:
				valueStrs[i] = "\"" + v + "\""
			default:
				return "", fmt.Errorf("type: %s, value: %v", reflect.TypeOf(v), v)
			}
		}
		sb.WriteString(strings.Join(valueStrs, ", "))
	}

	return sb.String(), nil
}
