package main

import (
	"flag"
	"goto/greenlight-m/internal/data"
	"goto/greenlight-m/internal/jsonlogger"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxOpenCons int
		maxIdleCons int
		maxIdleTime string
	}
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
}

type application struct {
	config config
	logger *jsonlogger.Logger
	models data.Models
}

func main() {
	logger := jsonlogger.New(os.Stdout, jsonlogger.LevelInfo)
	err := godotenv.Load()
	if err != nil {
		logger.LogError(err, nil)
	}
	var cfg config

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "string for db con establishment")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "30m", `maximum time duration for idle db connection`)
	flag.IntVar(&cfg.db.maxIdleCons, "max-idle-connections", 30, "maximum of allowed idle connections")
	flag.IntVar(&cfg.db.maxOpenCons, "max-open-connections", 30, "maximum of allowed open connections")
}
