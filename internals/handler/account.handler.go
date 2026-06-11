package handler

import (
	"strconv"

	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

var validate = validator.New()


type AccountHandler struct {
	Queries *db.Queries
	TokenMaker token.Maker
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
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
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

func (h *AccountHandler) GetAccount(c fiber.Ctx) error {
	id,err:=strconv.Atoi(c.Params("id"))

	payload := c.Locals("AuthorizationPayloadKey").(*token.Payload)


	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid account ID",
		})
	}
	if(id<=0){
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Account ID must be a positive integer",
		})
	}
	account,err:=h.Queries.GetAccount(c.Context(), int64(id))
	if account.Owner != payload.Username {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "You are not authorized to access this account",
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get account",
		})
	}

	return c.JSON(account)
}
type ListAccountsRequest struct {
	PageID int32 `query:"page_id" validate:"required,min=1"`
	PageSize int32 `query:"page_size" validate:"required,min=5,max=10"`
}

func (h *AccountHandler) ListAccounts(c fiber.Ctx) error {
	var req ListAccountsRequest
	if err := c.Bind().Query(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request query",
		})
	}
	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}
	accounts, err := h.Queries.ListAccounts(c.Context(), db.ListAccountsParams{
		Owner: c.Locals("AuthorizationPayloadKey").(*token.Payload).Username,
		Limit:   req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to list accounts",
		})
	}
	return c.JSON(accounts)
}