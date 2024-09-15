package repository

import (
	"context"

	"github.com/kiskislaya/finance-tracker/internal/entity"
)

type TransactionRepository interface {
	FindAll(ctx context.Context) ([]entity.Transaction, error)
	Save(ctx context.Context, entity entity.Transaction) (int64, error)
	Delete(ctx context.Context, id int64) error
}

type BalanceRepository interface {
	GetBalance(ctx context.Context) (int64, error)
}
