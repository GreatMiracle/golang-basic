package db

import (
	"context"
	"fmt"
)

type Store struct {
	db      DBTX
	Queries *Queries
}

// NewStore creates a new Store instance
func NewStore(db DBTX) *Store {
	return &Store{
		db:      db,
		Queries: New(db), // Initialize Queries
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TranferTx(ctx context.Context, params TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	errorExec := store.execTx(ctx, func(queries *Queries) error {
		var err error

		// Tạo bản ghi chuyển tiền
		result.Transfer, err = queries.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: params.FromAccountID,
			ToAccountID:   params.ToAccountID,
			Amount:        params.Amount})

		if err != nil {
			return err
		}

		// Tạo bản ghi: ghi nhận trừ tiền từ tài khoản 1
		result.FromEntry, err = queries.CreateEntry(ctx, CreateEntryParams{AccountID: params.FromAccountID, Amount: -params.Amount})
		if err != nil {
			return err
		}

		// Tạo bản ghi: ghi nhận cộng tiền vào tài khoản 2
		result.ToEntry, err = queries.CreateEntry(ctx, CreateEntryParams{AccountID: params.ToAccountID, Amount: params.Amount})
		if err != nil {
			return err
		}

		//TODO: update account's balance
		return nil
	})

	return result, errorExec

}
