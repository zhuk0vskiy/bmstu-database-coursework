package _interface

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IUserRepository
type IMonitorService interface {
	IncUsersOnline(request *dto.IncUsersOnlineRequest) error // Для получения данных пользователя
}
