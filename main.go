package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kcmikee/simplebank/api"
	db "github.com/kcmikee/simplebank/db/sqlc"

	// "github.com/kcmikee/simplebank/internal/env"
	"github.com/kcmikee/simplebank/util"
)

func main() {
	envConfig, err := util.LoadConfig(".")
	ctx := context.Background()
	cfg := config{
		addr: envConfig.ServerAddress,
		db: dbConfig{
			dsn: envConfig.DBSource,
		},
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Database
	conn, err := pgxpool.New(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(cfg.addr)
	if err != nil {
		log.Fatal("cannot start server:", err)
		panic(err)
	}
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}
