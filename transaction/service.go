package transaction

type service struct {
	repository Repository
}

type Service interface {
	CreateTransaction(input CreateTransactionInput) (Transaksi, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaksi, error) {
	transaction := Transaksi{}
	transaction.AcaraID = input.AcaraID
	transaction.UserID = input.UserID
	transaction.Total = input.Total
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}
