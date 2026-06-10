package handler

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/gofiber/fiber/v3"
)

type TransferHandler struct {
	Store *db.Store

}

type CreateTransferRequest struct {
	FromAccountID int64 `json:"from_account_id" validate:"required"`
	ToAccountID   int64 `json:"to_account_id" validate:"required"`
	Amount        int64 `json:"amount" validate:"required,gt=0"`
	Currency      string `json:"currency" validate:"required,oneof=USD EUR CAD"`
}

func (h *TransferHandler) CreateTransfer(c fiber.Ctx) error {
	var req CreateTransferRequest
	
	if err:=c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	valid, err := h.validAccount(c, req.FromAccountID, req.Currency);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to validate from account",
		})
	}	else if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "From account does not exist or currency mismatch",
		})
	}
	valid, err = h.validAccount(c, req.ToAccountID, req.Currency);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to validate to account",
		})
	}	else if !valid {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "To account does not exist or currency mismatch",
		})
	}

	arg:= db.CreateTransferParams{
		FromAccountID: req.FromAccountID,
		ToAccountID: req.ToAccountID,
		Amount: req.Amount,
	}
	result,err:=h.Store.TransferTx(c.Context(), arg)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create transfer",
		})
	}
	return c.Status(fiber.StatusOK).JSON(result)

}

func (h *TransferHandler) validAccount(c fiber.Ctx, accountID int64, currency string) (bool, error){

	account,err:=h.Store.GetAccount(c.Context(), accountID)
	if err != nil {
		return false, err
	}
	if account.Currency != currency {
		return false, nil
	}
	return true, nil
} 