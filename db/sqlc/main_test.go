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

const defaultTestDBSource = "postgresql://root:" + "secret@localhost:5433/simple_bank?sslmode=disable"

func getTestDBSource() string {
	if dbSource := os.Getenv("DB_SOURCE"); dbSource != "" {
		return dbSource
	}

	config, err := util.LoadConfig("../..")
	if err == nil && config.DBSource != "" {
		return config.DBSource
	}

	return defaultTestDBSource
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
