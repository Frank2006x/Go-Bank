package main

import (
	"context"
	"fmt"
	"log"

	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals"
	"github.com/Frank2006x/simple-bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool
var Queries *db.Queries
var Store *db.Store



func main() {
	config,err:=util.LoadConfig("./")
	fmt.Println(config.DBSource)
	fmt.Println(config.ServerAddress)
	if err!=nil{
		log.Fatal("cannot load config:", err)
	}
	
	DB, err := pgxpool.New(
		context.Background(),
		config.DBSource,
	)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	Queries = db.New(DB)
	Store = db.NewStore(DB)

	Server := internals.NewServer(Store)
	err = Server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}