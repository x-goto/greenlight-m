package main

import (
	"context"
	"flag"
	"goto/greenlight-m/internal/data"
	"goto/greenlight-m/pkg/client/pgsql"
	"goto/greenlight-m/pkg/httpcodec"
	"goto/greenlight-m/pkg/httpcodec/json"
	"goto/greenlight-m/pkg/jsonlogger"
	"net/http"
	"os"
	"strconv"
	"time"

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
	config       config
	logger       *jsonlogger.Logger
	codec        httpcodec.HTTPCodec
	repositories data.Repositories
	server       *http.Server
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

	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("PSQL_DB_DSN"), "string for db con establishment")
	flag.StringVar(&cfg.db.maxIdleTime, "db-max-idle-time", "30m", `maximum time duration for idle db connection`)
	flag.IntVar(&cfg.db.maxIdleCons, "max-idle-connections", 30, "maximum of allowed idle connections")
	flag.IntVar(&cfg.db.maxOpenCons, "max-open-connections", 30, "maximum of allowed open connections")

	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db, err := pgsql.NewClient(ctx, cfg.db.dsn, cfg.db.maxOpenCons, cfg.db.maxIdleCons, cfg.db.maxIdleTime)
	if err != nil {
		logger.LogFatal(err, nil)
	}

	defer db.Close()

	logger.LogInfo("DB connection pool established", nil)

	app := &application{
		config:       cfg,
		logger:       logger,
		codec:        json.New(),
		repositories: data.NewPQRepositories(db),
		//server will be initialized in serve() ***server.go file
	}

	if err = app.serve(); err != nil {
		logger.LogFatal(err, nil)
	}

}
