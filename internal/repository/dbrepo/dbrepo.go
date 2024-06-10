package dbrepo

import (
	"database/sql"

	"github.com/saim61/udemy_bed_and_breakfast/internal/config"
	"github.com/saim61/udemy_bed_and_breakfast/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}