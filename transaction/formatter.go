package transaction

type TransactionFormatter struct {
	ID      int    `json:"id"`
	AcaraID int    `json:"acara_id"`
	UserID  int    `json:"user_id"`
	Total   int    `json:"total"`
	Status  string `json:"status"`
}

func FormatTransaction(transaction Transaksi) TransactionFormatter {
	formatter := TransactionFormatter{}
	formatter.ID = transaction.ID
	formatter.AcaraID = transaction.AcaraID
	formatter.UserID = transaction.UserID
	formatter.Total = transaction.Total
	formatter.Status = transaction.Status
	return formatter
}
