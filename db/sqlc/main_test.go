package db

import (
	"database/sql"
	"os"
	"testing"

	"github.com/WuzorGiftKnowledge/SimpleBank/util"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

// const (
// 	driver     = "postgres"
// )
// var dataSource = "postgresql://root:root@localhost/simplebank?sslmode=disable"

var testQueries *Queries
var testDB *sql.DB
var testStore *Store

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal().Err(err).Msg("unable to load config")
	}

	testDB, err = sql.Open(config.DRIVER, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to connect to database")
	}
	testQueries = New(testDB)
	testStore = NewStore(testDB)
	os.Exit(m.Run())
}
