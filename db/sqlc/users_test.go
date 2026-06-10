package db

import (
	"context"
	"testing"
	"time"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := CreateUserParams{
		Username:     util.RandomOwner(),
		Email:        util.RandomEmail(),
		PasswordHash: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PasswordHash, user.PasswordHash)

	require.True(t, user.PasswordChangedAt.Time.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.PasswordHash, user2.PasswordHash)
	require.WithinDuration(t, user1.PasswordChangedAt.Time, user2.PasswordChangedAt.Time, time.Second)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	newEmail := util.RandomEmail()
	newPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := UpdateUserParams{
		Username: user1.Username,
		Email: pgtype.Text{
			String: newEmail,
			Valid:  true,
		},
		PasswordHash: pgtype.Text{
			String: newPassword,
			Valid:  true,
		},
		PasswordChangedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.Email.String, user2.Email)
	require.Equal(t, arg.PasswordHash.String, user2.PasswordHash)
	require.WithinDuration(t, arg.PasswordChangedAt.Time, user2.PasswordChangedAt.Time, time.Second)
}
