# Трекер финансов

## Архитектура

### HTTP API

- Получение баланса: `GET /v1/balance`
- Получение списка транзакций: `GET /v1/transactions`
- Добавление транзакции списания: `POST /v1/withdraw`
- Добавление транзакции пополнения: `POST /v1/deposit`
- Удаление транзакции: `DELETE /v1/transactions/{transactionId}`

#### Контракты

```go
type V1GetBalanceResponse struct {
	Balance string `json:"balance"`
}

type V1Transaction struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type V1ListTransactionsResponse struct {
	Transactions []V1Transaction `json:"transactions"`
}

type V1WithdrawRequest struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}

type V1DepositRequest struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
```

#### Структура базы данных

```go
type Transaction struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Amount    int64     `db:"amount"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
}
```
