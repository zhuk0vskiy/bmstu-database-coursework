package _interface

import (
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model/dto"
)

type IRoomService interface {
	Get(request *dto.GetRoomRequest) (*model.Room, error)                   // Для отдельного вывода изначальной информации на странице для отдельной комнаты при обновлении
	GetByStudio(request *dto.GetRoomByStudioRequest) ([]*model.Room, error) // Для поиска незаброненных комнат студии и при изменении комнат
	Add(request *dto.AddRoomRequest) error                                  // Для добавления
	Update(request *dto.UpdateRoomRequest) error                            // Для обоновления
	Delete(request *dto.DeleteRoomRequest) error                            // Для удаления
}
