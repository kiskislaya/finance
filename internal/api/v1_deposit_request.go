package api

type V1DepositRequest struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
