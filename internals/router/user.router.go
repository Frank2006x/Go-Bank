package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupUserRouter(app *fiber.App, store *db.Store) {
	userHandler := &handler.UserHandler{
		Queries: store.Queries,
	}

	userGroup := app.Group("/users")
	userGroup.Post("/", userHandler.CreateUser)
}
