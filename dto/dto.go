package dto

type SignIn struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
