package data

import (
	"database/sql"
	"goto/greenlight-m/internal/data/token"
	"goto/greenlight-m/internal/data/token/tokendbs"
	"goto/greenlight-m/internal/data/user"
	"goto/greenlight-m/internal/data/user/userdbs"
)

type Repositories struct {
	Users  user.Repository
	Tokens token.Repository
}

func NewPQRepositories(db *sql.DB) Repositories {
	return Repositories{
		Users:  &userdbs.PQUserRepository{DB: db},
		Tokens: &tokendbs.PQTokenRepository{DB: db},
	}
}
