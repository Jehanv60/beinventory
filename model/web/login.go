package web

type LoginRequest struct {
	Pengguna string `validate:"required,max=100,min=1" json:"pengguna"`
	Email    string `validate:"required" json:"email"`
	Sandi    string `validate:"required,max=100,min=1" json:"sandi"`
}
