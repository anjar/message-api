package services

type InputPagination struct {
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	OrderBy string `json:"order_by"`
}