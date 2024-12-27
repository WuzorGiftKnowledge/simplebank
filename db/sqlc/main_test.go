package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

const (
	driver     = "postgres"
)
var dataSource = "postgresql://root:root@localhost/simplebank?sslmode=disable"

var testQueries *Queries
var testDB *sql.DB
var testStore *Store

func TestMain(m *testing.M) {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	connString := dataSource
	log.Printf("host=%s, port=%s, dbName=%s, username=%s", dbHost,dbPort, dbName, dbUser)
	if dbHost != "" && dbPort != "" && dbName != "" {
		connString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
	}
	log.Print(connString)
	var err error
	testDB, err = sql.Open(driver, connString)
	if err != nil {
		log.Fatal("unable to connect to database:", err)
	}
	testQueries = New(testDB)
	testStore = NewStore(testDB)
	os.Exit(m.Run())
}
