package service

import (
	"context"

	"github.com/kiskislaya/finance-tracker/internal/model"
)

type TransactionService interface {
	GetAll(ctx context.Context) ([]model.Transaction, error)
	Deposit(ctx context.Context, name string, amount int64) error
	Withdraw(ctx context.Context, name string, amount int64) error
	Delete(ctx context.Context, id int64) error
}

type BalanceService interface {
	GetBalance(ctx context.Context) (int64, error)
}
