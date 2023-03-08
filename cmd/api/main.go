package main

import (
	"flag"
	"goto/greenlight-m/internal/data"
	"goto/greenlight-m/internal/jsonlogger"
	"os"
	"strconv"

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

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.LogError(err, nil)
	}

	flag.StringVar(&cfg.env, "env", os.Getenv("ENV"), "ENV")
	flag.IntVar(&cfg.port, "port", port, "port on which application will run")

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "string for db con establishment")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "30m", `maximum time duration for idle db connection`)
	flag.IntVar(&cfg.db.maxIdleCons, "max-idle-connections", 30, "maximum of allowed idle connections")
	flag.IntVar(&cfg.db.maxOpenCons, "max-open-connections", 30, "maximum of allowed open connections")
}
