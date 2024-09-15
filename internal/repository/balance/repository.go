package balance

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kiskislaya/finance-tracker/internal/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.BalanceRepository {
	return repo{
		db: db,
	}
}

func (r repo) GetBalance(ctx context.Context) (int64, error) {
	query := `
	SELECT
		SUM(CASE WHEN type = 'DEPOSIT' THEN amount ELSE 0 END) -
		SUM(CASE WHEN type = 'WITHDRAW' THEN amount ELSE 0 END) AS balance
	FROM transactions;`

	row := r.db.QueryRow(ctx, query)

	var balance int64
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}
