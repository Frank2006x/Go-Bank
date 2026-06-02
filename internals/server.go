package internals

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/router"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Server struct {
	Store *db.Store
	Router *fiber.App
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		Store: store,
	}
	server.Router = fiber.New()
	server.Router.Use(logger.New())
	router.SetupAccountRouter(server.Router, server.Store)
	return server
}

func (server *Server) Start(address string) error {
	return server.Router.Listen(address)
}