package internals

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/internals/router"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/Frank2006x/simple-bank/util"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

type Server struct {
	Store *db.Store
	Router *fiber.App
	tokenMaker token.Maker
	config util.Config
}

func NewServer(config util.Config, store *db.Store) (*Server ,error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, err
	}
	server := &Server{
		Store: store,
		tokenMaker: tokenMaker,
		config: config,
	}
	server.Router = fiber.New()
	server.Router.Use(logger.New())
	router.SetupAccountRouter(server.Router, server.Store, server.tokenMaker)
	router.SetupUserRouter(server.Router, server.Store, server.tokenMaker, server.config)
	router.SetupTransferRouter(server.Router, server.Store,server.tokenMaker)
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.Router.Listen(address)
}