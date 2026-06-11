package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/Frank2006x/simple-bank/internals/middleware"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/gofiber/fiber/v3"
)

func SetupTransferRouter(app *fiber.App, store *db.Store, tokenMaker token.Maker) {
	transferHandler := &handler.TransferHandler{
		Store: store,
		TokenMaker: tokenMaker,
	}

	transferGroup := app.Group("/transfers")
	transferGroup.Use(middleware.AuthMiddleware(tokenMaker))
	transferGroup.Post("/", transferHandler.CreateTransfer)
	
}