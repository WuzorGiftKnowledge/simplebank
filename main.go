package main

import (
	"database/sql"
	"os"

	"github.com/WuzorGiftKnowledge/SimpleBank/api"
	db "github.com/WuzorGiftKnowledge/SimpleBank/db/sqlc"
	"github.com/WuzorGiftKnowledge/SimpleBank/util"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DRIVER, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server:")
	}

}
