package user

// register
type RegisterUserInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Nama     string `json:"nama" binding:"required"`
}

// login

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}

//cek email input
type CheckEmailInput struct {
	Email string `json:"email" form:"email" binding:"required,email"`
}
