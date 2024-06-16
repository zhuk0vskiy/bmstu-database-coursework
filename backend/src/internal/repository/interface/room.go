package _interface

import (
	"context"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model/dto"
)

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=IRoomRepository
type IRoomRepository interface {
	Get(ctx context.Context, request *dto.GetRoomRequest) (*model.Room, error)                   // Для отдельного вывода изначальной информации на странице для отдельной комнаты при обновлении
	GetByStudio(ctx context.Context, request *dto.GetRoomByStudioRequest) ([]*model.Room, error) // Для поиска незаброненных комнат студии и при изменении комнат
	Add(ctx context.Context, request *dto.AddRoomRequest) error                                  // Для добавления //TODO: Добавить dto
	Update(ctx context.Context, request *dto.UpdateRoomRequest) error                            // Для обоновления
	Delete(ctx context.Context, request *dto.DeleteRoomRequest) error                            // Для удаления
}
