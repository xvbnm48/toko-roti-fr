package postgre

import "database/sql"

type DbReadWriter struct {
	DB *sql.DB
}

func NewDatabaseWriter() *DatabaseWriter {
	return &DatabaseWriter{}
}
