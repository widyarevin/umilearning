package transaction

type CreateTransactionInput struct {
	UserID  int `json:"user_id" binding:"required"`
	AcaraID int `json:"acara_id" binding:"required"`
	Total   int `json:"total" binding:"required"`
}
