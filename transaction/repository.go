package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	Save(transaction Transaksi) (Transaksi, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) Save(transaksi Transaksi) (Transaksi, error) {
	err := r.db.Create(&transaksi).Error

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
}
