package handler

import (
	db "github.com/Frank2006x/simple-bank/db/sqlc"
	"github.com/Frank2006x/simple-bank/token"
	"github.com/Frank2006x/simple-bank/util"
	"github.com/gofiber/fiber/v3"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserHandler struct {
	Queries *db.Queries
	TokenMaker token.Maker
	Config     util.Config
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

type loginUserRequest struct {
	Username string `json:"username" validate:"required,alphanum"`
	Password string `json:"password" validate:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	User        userResponse `json:"user"`
}


func (h* UserHandler) LoginUser(c fiber.Ctx) error {
	var req loginUserRequest
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

	user, err := h.Queries.GetUser(c.Context(), req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get user: " + err.Error(),
		})
	}

	isMatch := util.CheckPasswordHash(req.Password, user.PasswordHash)
	if !isMatch {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	accessToken, err := h.TokenMaker.CreateToken(user.Username, h.Config.TokenDuration)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create access token: " + err.Error(),
		})
	}
	response := loginUserResponse{
		AccessToken: accessToken,
		User: newUserResponse(user),
	}
	return c.Status(fiber.StatusOK).JSON(response)
}