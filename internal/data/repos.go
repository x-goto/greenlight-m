package data

import (
	"database/sql"
	"goto/greenlight-m/internal/data/users"
	"goto/greenlight-m/internal/data/users/dbs"
)

type Repositories struct {
	Users users.Repository
}

func NewPQRepositories(db *sql.DB) Repositories {
	return Repositories{
		Users: &dbs.PQUserRepository{DB: db},
	}
}
