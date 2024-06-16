package _interface

import (
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model/dto"
)

type IAuthService interface {
	SignIn(request *dto.SignInRequest) error
	LogIn(request *dto.LogInRequest) (string, error)
}
