package repository

import (
	"database/sql"
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

	var sb strings.Builder
	var dest string
	for rows.Next() {
		if err = rows.Scan(&dest); err != nil {
			return "", err
		}
		sb.WriteString(dest)
	}

	return sb.String(), nil
}
