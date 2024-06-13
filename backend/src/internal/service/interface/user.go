package _interface

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
)

type IUserService interface {
	Get(request *dto.GetUserRequest) (*model.User, error)                      // Для вывода пользователю его данных
	Update(request *dto.UpdateUserRequest) error                               // Для обновления данных пользователя
	Delete(request *dto.DeleteUserRequest) error                               // Для удаления пользователя
	GetReserves(request *dto.GetUserReservesRequest) ([]*model.Reserve, error) // Для вывода пользователю его броней
	GetByLogin(request *dto.GetUserByLoginRequest) (*model.User, error)
	//GetAll(id UserId) ([]*User, error)
}
