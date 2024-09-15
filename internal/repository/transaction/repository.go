package transaction

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kiskislaya/finance-tracker/internal/entity"
	"github.com/kiskislaya/finance-tracker/internal/repository"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.TransactionRepository {
	return repo{
		db: db,
	}
}

func (r repo) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	query := `select (id, name, amount, type, created_at) from transactions`

	rows, err := r.db.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	var entities []entity.Transaction

	var ent entity.Transaction
	for rows.Next() {
		err = rows.Scan(&ent)
		if err != nil {
			return nil, err
		}
		entities = append(entities, ent)
	}

	return entities, nil
}

func (r repo) Save(ctx context.Context, ent entity.Transaction) (int64, error) {
	query := `insert into transactions (name, amount, type) values ($1, $2, $3) returning id`

	row := r.db.QueryRow(
		ctx,
		query,
		ent.Name,
		ent.Amount,
		ent.Type,
	)

	var result = entity.Transaction{}
	err := row.Scan(&result.ID)
	if err != nil {
		return 0, err
	}

	return result.ID, nil
}

func (r repo) Delete(ctx context.Context, id int64) error {
	query := `delete from transactions where id = $1`

	_, err := r.db.Exec(
		ctx,
		query,
		id,
	)

	if err != nil {
		return err
	}

	return nil
}
