package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupAccountRouter(app *fiber.App, store *db.Store) {
	accountHandler := &handler.AccountHandler{
		Queries: store.Queries,
	}

	accountGroup := app.Group("/accounts")
	accountGroup.Post("/", accountHandler.CreateAccount)
}