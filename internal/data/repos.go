package data

import (
	"database/sql"
	"goto/greenlight-m/internal/data/user"
	"goto/greenlight-m/internal/data/user/dbs"
)

type Repositories struct {
	Users user.Repository
}

func NewPQRepositories(db *sql.DB) Repositories {
	return Repositories{
		Users: &dbs.PQUserRepository{DB: db},
	}
}
