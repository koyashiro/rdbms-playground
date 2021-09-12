package repository

import (
	"database/sql"
	"strconv"
	"strings"
)

type DBRepository interface {
	Execute(port int, query string) (string, error)
}

type PostgresRepositoryImpl struct{}

func NewPostgresRepository() DBRepository {
	return &PostgresRepositoryImpl{}
}

func (r *PostgresRepositoryImpl) Execute(port int, query string) (string, error) {
	db, err := sql.Open("postgres", "user=postgres password=password dbname=postgres sslmode=disable port="+strconv.Itoa(port))
	if err != nil {
		return "", err
	}

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
