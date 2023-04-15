package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Nama:     user.Nama,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Token:    token,
	}
	return formatter
}
