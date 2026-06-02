package db

import (
	"context"
	"testing"

	"github.com/Frank2006x/simple-bank/util"
	"github.com/stretchr/testify/require"
)
func createRandomEntry(t *testing.T, account Account) Entry {
	arg := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}
	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account)	

}

func TestGetEntryByAccountID(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}
	arg := GetEntryByAccountIDParams{
		AccountID: account.ID,
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.GetEntryByAccountID(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

}

func TestHistoryEntry(t *testing.T) {
	for i := 0; i < 10; i++ {
		account := createRandomAccount(t)
		createRandomEntry(t, account)
	}
	arg := HistoryParams{
		Limit: 5,
		Offset: 5,
	}
	entries, err := testQueries.History(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)
	
	

}