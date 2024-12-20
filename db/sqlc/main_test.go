package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_"github.com/lib/pq"
)
const (
	driver = "postgres"
	dataSource= "postgresql://root:root@127.0.0.1:5432/simplebank?sslmode=disable"
)
var testQueries *Queries

func TestMain(m *testing.M) {
conn, err:= sql.Open(driver, dataSource)
if err != nil{
	log.Fatal("unable to connect to database:", err)
}
  testQueries = New(conn)
 os.Exit(m.Run())
}