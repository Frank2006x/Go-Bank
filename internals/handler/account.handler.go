package handler

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/gofiber/fiber/v3"
)

type AccountHandler struct {
	Queries *db.Queries
}

type CreateAccountRequest struct {
	Owner    string `json:"owner" validate:"required"`
	Currency string `json:"currency" validate:"required,oneof=USD EUR CAD"`
}

func (h *AccountHandler) CreateAccount(c fiber.Ctx) error {
	var req CreateAccountRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	account,err:=h.Queries.CreateAccount(c.Context(), db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create account",
		})
	}
	return c.Status(fiber.StatusCreated).JSON(account)
}
