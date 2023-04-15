package acara

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Acara, error)
	FindByID(ID int) (Acara, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Acara, error) {
	var acaras []Acara

	err := r.db.Find(&acaras).Error
	if err != nil {
		return acaras, err
	}

	return acaras, nil
}

func (r *repository) FindByID(ID int) (Acara, error) {
	var acara Acara

	err := r.db.Where("id = ?", ID).Find(&acara).Error

	if err != nil {
		return acara, err
	}

	return acara, nil
}
