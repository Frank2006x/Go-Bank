package db

import (
	"context"
	"testing"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Amount: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}


func TestGetTransferByFromAccountID(t *testing.T) {

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID: account2.ID,
			Amount: 10,
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
	}
	arg := GetTransferByFromAccountIDParams{
		FromAccountID: account1.ID,
		Limit: 5,
	}
	transfers, err := testQueries.GetTransferByFromAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.Equal(t, account1.ID, transfer.FromAccountID)
	}
}
	
func TestGetTransferByToAccountID(t *testing.T) {

	account1 := createRandomAccount(t)	
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID: account2.ID,
			Amount: 10,
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
	}
	arg := GetTransferByToAccountIDParams{
		ToAccountID: account2.ID,
		Limit: 5,
	}
	transfers, err := testQueries.GetTransferByToAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.Equal(t, account2.ID, transfer.ToAccountID)
	}
}

func TestListTransfer(t *testing.T) {

	account1 := createRandomAccount(t)	
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		arg := CreateTransferParams{
			FromAccountID: account1.ID,
			ToAccountID: account2.ID,
			Amount: 10,
		}
		_, err := testQueries.CreateTransfer(context.Background(), arg)
		require.NoError(t, err)
	}
	arg := ListTransfersBetweenAccountsParams{
		FromAccountID: account1.ID,
		ToAccountID: account2.ID,
		Limit: 5,
	}
	transfers, err := testQueries.ListTransfersBetweenAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfers)
	require.Len(t, transfers, 5)
	for _, transfer := range transfers {
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
	}
}