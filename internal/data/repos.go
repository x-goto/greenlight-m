package data

import (
	"database/sql"
	"goto/greenlight-m/internal/data/users"
)

type Repositories struct {
	Users users.Repository
}

func NewPQRepositories(db *sql.DB) Repositories {
	return Repositories{}
}
