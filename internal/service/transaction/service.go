package transaction

import (
	"context"

	"github.com/kiskislaya/finance-tracker/internal/entity"
	"github.com/kiskislaya/finance-tracker/internal/model"
	"github.com/kiskislaya/finance-tracker/internal/repository"
	"github.com/kiskislaya/finance-tracker/internal/service"
	"github.com/redis/go-redis/v9"
)

type serv struct {
	transactionRepository repository.TransactionRepository
	redisClient           *redis.Client
}

func NewService(transactionRepository repository.TransactionRepository, redisClient *redis.Client) service.TransactionService {
	return serv{
		transactionRepository: transactionRepository,
		redisClient:           redisClient,
	}
}

func (s serv) GetAll(ctx context.Context) ([]model.Transaction, error) {
	entities, err := s.transactionRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var models []model.Transaction
	for _, entity := range entities {
		models = append(models, model.Transaction{
			ID:     entity.ID,
			Name:   entity.Name,
			Amount: entity.Amount,
			Type:   entity.Type,
		})
	}

	return models, nil
}

func (s serv) Deposit(ctx context.Context, name string, amount int64) error {
	_, err := s.transactionRepository.Save(ctx, entity.Transaction{
		Name:   name,
		Amount: amount,
		Type:   "DEPOSIT",
	})
	if err != nil {
		return err
	}

	if err := s.redisClient.Del(ctx, "balance").Err(); err != nil {
		return err
	}

	return nil
}

func (s serv) Withdraw(ctx context.Context, name string, amount int64) error {
	_, err := s.transactionRepository.Save(ctx, entity.Transaction{
		Name:   name,
		Amount: amount,
		Type:   "WITHDRAW",
	})
	if err != nil {
		return err
	}

	if err := s.redisClient.Del(ctx, "balance").Err(); err != nil {
		return err
	}

	return nil
}

func (s serv) Delete(ctx context.Context, id int64) error {
	err := s.transactionRepository.Delete(ctx, id)
	if err != nil {
		return err
	}

	if err := s.redisClient.Del(ctx, "balance").Err(); err != nil {
		return err
	}

	return nil
}
