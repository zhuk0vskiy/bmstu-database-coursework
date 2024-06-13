package _interface

import (
	"app/src/internal/model/dto"
)

type IAuthService interface {
	SignIn(request *dto.SignInRequest) error
	LogIn(request *dto.LogInRequest) (string, error)
}
