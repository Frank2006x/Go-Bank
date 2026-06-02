package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type store struct {
	*Queries
	db *pgxpool.Pool
}


func NewStore(db *pgxpool.Pool) *store {
	return &store{
		Queries: New(db),
		db: db,
	}
}


func (s *store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := s.db.BeginTx(ctx, pgx.TxOptions{})

	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := s.WithTx(tx)
	if err := fn(q); err != nil {
		return err
	}

	return tx.Commit(ctx)

}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
}



func (s *store) TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}
		
		if(result.FromAccount.ID < result.ToAccount.ID){
			 result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Balance: -arg.Amount,
			})

			if err != nil {
				return err
			}
			result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Balance:  arg.Amount,
			})
			if err != nil {
				return err
			}

		}else{
			 result.ToAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.ToAccountID,
				Balance:  arg.Amount,
			})
			if err != nil {
				return err
			}
			result.FromAccount, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
				ID:     arg.FromAccountID,
				Balance: -arg.Amount,
			})
			if err != nil {
				return err
			}
		}
		
		
		return nil
	})
	return result, err
}