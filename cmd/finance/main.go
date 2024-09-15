package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/kiskislaya/finance-tracker/internal/api"
	balanceRepository "github.com/kiskislaya/finance-tracker/internal/repository/balance"
	transactionRepository "github.com/kiskislaya/finance-tracker/internal/repository/transaction"
	balanceService "github.com/kiskislaya/finance-tracker/internal/service/balance"
	transactionService "github.com/kiskislaya/finance-tracker/internal/service/transaction"
	"github.com/redis/go-redis/v9"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, "postgres://finance:finance@localhost:5433/finance")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	redisClient := redis.NewClient(&redis.Options{
		Addr: net.JoinHostPort("localhost", "6380"),
	})

	transactionService := transactionService.NewService(
		transactionRepository.NewRepository(dbpool),
		redisClient,
	)

	balanceService := balanceService.NewService(
		balanceRepository.NewRepository(dbpool),
		redisClient,
	)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /v1/balance", func(w http.ResponseWriter, r *http.Request) {
		balance, err := balanceService.GetBalance(ctx)
		if err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		balanceStr := strconv.Itoa(int(balance))

		balanceResponse := api.V1GetBalanceResponse{
			Balance: balanceStr,
		}
		api.WriteJSONResponse(w, http.StatusOK, balanceResponse)
	})

	mux.HandleFunc("GET /v1/transactions", func(w http.ResponseWriter, r *http.Request) {
		transactionsResponse := api.V1ListTransactionsResponse{
			Transactions: []api.V1Transaction{},
		}

		transactionsModels, err := transactionService.GetAll(ctx)
		if err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		for _, tm := range transactionsModels {
			id := strconv.Itoa(int(tm.ID))
			amount := strconv.Itoa(int(tm.Amount))

			transactionsResponse.Transactions = append(transactionsResponse.Transactions, api.V1Transaction{
				ID:     id,
				Name:   tm.Name,
				Amount: amount,
				Type:   tm.Type,
			})
		}

		api.WriteJSONResponse(w, http.StatusOK, transactionsResponse)
	})

	mux.HandleFunc("POST /v1/withdraw", func(w http.ResponseWriter, r *http.Request) {
		var withdrawRequest api.V1WithdrawRequest

		dec := json.NewDecoder(r.Body)
		dec.Decode(&withdrawRequest)

		amount, err := strconv.Atoi(withdrawRequest.Amount)
		if err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		if err := transactionService.Withdraw(ctx, withdrawRequest.Name, int64(amount)); err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("POST /v1/deposit", func(w http.ResponseWriter, r *http.Request) {
		var depositRequest api.V1DepositRequest

		dec := json.NewDecoder(r.Body)
		dec.Decode(&depositRequest)

		amount, err := strconv.Atoi(depositRequest.Amount)
		if err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		if err := transactionService.Deposit(ctx, depositRequest.Name, int64(amount)); err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	mux.HandleFunc("DELETE /v1/transactions/{transactionId}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("transactionId"))
		if err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}

		if err := transactionService.Delete(ctx, int64(id)); err != nil {
			api.WriteProblemResponse(w, api.ProblemDetails{
				Status: http.StatusInternalServerError,
			})
			return
		}
	})

	httpServer := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}
