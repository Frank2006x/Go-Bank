package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool
var testStore *Store

func getTestDBSource() string {
	if dbSource := os.Getenv("DB_SOURCE"); dbSource != "" {
		return dbSource
	}

	config, err := util.LoadConfig("../..")
	if err == nil && config.DBSource != "" {
		return config.DBSource
	}

	log.Fatal("cannot load test db source from DB_SOURCE or config")
	return ""
}

func TestMain(m *testing.M) {
	dbSource := getTestDBSource()

	var err error
	testDB, err = pgxpool.New(
		context.Background(),
		dbSource,
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)
	testStore = NewStore(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)
}
