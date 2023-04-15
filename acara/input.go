package acara

type GetAcaraDetailInput struct {
	ID int `uri:"id" binding:"required"`
}
