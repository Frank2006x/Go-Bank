package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/gofiber/fiber/v3"
)

func SetupTransferRouter(app *fiber.App, store *db.Store) {
	transferHandler := &handler.TransferHandler{
		Store: store,
	}

	transferGroup := app.Group("/transfers")
	transferGroup.Post("/", transferHandler.CreateTransfer)
	
}