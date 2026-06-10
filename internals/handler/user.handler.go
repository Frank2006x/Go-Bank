package handler

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/util"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct {
	Queries *db.Queries
}

type CreateUserRequest struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

type userResponse struct {
	Username          string             `json:"username"`
	Email             string             `json:"email"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data: " + err.Error(),
		})
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	arg := db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: hashedPassword,
		Email:        req.Email,
	}

	user, err := h.Queries.CreateUser(c.Context(), arg)
	if err != nil {
		if err == pgx.ErrNoRows {
			// This shouldn't happen with INSERT RETURNING but good practice
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user: " + err.Error(),
		})
	}

	rsp := newUserResponse(user)
	return c.Status(fiber.StatusCreated).JSON(rsp)
}
