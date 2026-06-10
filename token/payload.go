package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string `json:"username"`
	IssuedAt  int64  `json:"issued_at"`
	ExpiredAt int64  `json:"expired_at"`
}

var (
	ErrExpiredToken = errors.New("token has expired")
)

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now().Unix(),
		ExpiredAt: time.Now().Add(duration).Unix(),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().Unix() > payload.ExpiredAt {
		return ErrExpiredToken
	}

	return nil
}