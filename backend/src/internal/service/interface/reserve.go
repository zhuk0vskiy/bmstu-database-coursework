package _interface

import (
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model"
	"github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model/dto"
)

type IReserveService interface {
	Add(request *dto.AddReserveRequest) error       // Для создания брони
	Delete(request *dto.DeleteReserveRequest) error // Для удаления брони
	GetAll(request *dto.GetAllReserveRequest) (equipments []*model.Reserve, err error)
	//Get(id int64) (*Reserve, error)
	//GetNotReservedRooms(startTime, endTime time.Time, studioId int64) ([]*Room, error)
	//GetNotReservedProducers(startTime time.Time, endTime time.Time, studioId int64) ([]*Producer, error)
	//GetNotReservedEquipments(startTime time.Time, endTime time.Time, studioId int64) ([]*Equipment, error)
	//GetNotReservedInstrumentalists(startTime time.Time, endTime time.Time, studioId int64) ([]*Instrumentalist, error)
}
