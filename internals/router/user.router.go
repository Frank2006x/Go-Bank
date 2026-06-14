package router

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/handler"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/Frank2006x/simple-bank/util"
	"github.com/gofiber/fiber/v3"
)

func SetupUserRouter(router *fiber.App, store *db.Store, tokenMaker token.Maker, config util.Config) {
	userHandler := &handler.UserHandler{
		Queries: store.Queries,
		TokenMaker: tokenMaker,
		Config: config,
	}

	userGroup := router.Group("/users")
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Post("/login", userHandler.LoginUser)
}
