package dbs

import (
	"database/sql"
)

type PQUserRepository struct {
	DB *sql.DB
}
