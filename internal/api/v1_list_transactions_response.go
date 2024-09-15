package api

type V1Transaction struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount string `json:"amount"`
	Type   string `json:"type"`
}

type V1ListTransactionsResponse struct {
	Transactions []V1Transaction `json:"transactions"`
}
