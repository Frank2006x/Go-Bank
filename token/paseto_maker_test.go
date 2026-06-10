package token

import (
	"testing"
	"time"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	username := util.RandomOwner()
	duration :=	time.Minute // 1 minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)


	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)


	require.Equal(t, username, payload.Username)
	require.WithinDuration(t, issuedAt, time.Unix(payload.IssuedAt, 0), time.Second)
	require.WithinDuration(t, expiredAt, time.Unix(payload.ExpiredAt, 0), time.Second)

}