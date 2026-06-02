package main

import (
	"context"
	"log"

	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var Queries *db.Queries
var Store *db.Store

const (
	dbSource = "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	DB, err := pgxpool.New(
		context.Background(),
		dbSource,
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	Queries = db.New(DB)
	Store = db.NewStore(DB)

	Server := internals.NewServer(Store)
	err = Server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}