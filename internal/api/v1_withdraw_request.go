package api

type V1WithdrawRequest struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
}
