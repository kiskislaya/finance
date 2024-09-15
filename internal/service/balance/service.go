package balance

import (
	"context"
	"time"

	"github.com/kiskislaya/finance-tracker/internal/repository"
	"github.com/kiskislaya/finance-tracker/internal/service"
	"github.com/redis/go-redis/v9"
)

type serv struct {
	balanceRepository repository.BalanceRepository
	redisClient       *redis.Client
}

func NewService(balanceRepository repository.BalanceRepository, redisClient *redis.Client) service.BalanceService {
	return serv{
		balanceRepository: balanceRepository,
		redisClient:       redisClient,
	}
}

func (s serv) GetBalance(ctx context.Context) (int64, error) {
	val := s.redisClient.Get(ctx, "balance")

	cachedBalance, err := val.Int64()
	if err == redis.Nil {
		balance, err := s.balanceRepository.GetBalance(ctx)
		if err != nil {
			return 0, err
		}
		s.redisClient.SetEx(ctx, "balance", balance, time.Hour)
		return balance, nil
	} else if err != nil {
		return 0, err
	}

	return cachedBalance, nil
}
