package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/Frank2006x/simple-bank/internals/middleware"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/gofiber/fiber/v3"
)

func SetupAccountRouter(app *fiber.App, store *db.Store, tokenMaker token.Maker) {
	accountHandler := &handler.AccountHandler{
		Queries: store.Queries,
		TokenMaker: tokenMaker,
	}

	accountGroup := app.Group("/accounts")
	accountGroup.Use(middleware.AuthMiddleware(tokenMaker))
	accountGroup.Post("/", accountHandler.CreateAccount)
	accountGroup.Get("/:id", accountHandler.GetAccount)
	accountGroup.Get("/", accountHandler.ListAccounts)
}