package impl

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
	repositoryInterface "app/src/internal/repository/interface"
	serviceInterface "app/src/internal/service/interface"
	"context"
	"fmt"
)

type ReserveService struct {
	reserveRepo repositoryInterface.IReserveRepository
}

func NewReserveService(
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IReserveService {
	return &ReserveService{
		reserveRepo: reserveRepo,
	}
}

func (s ReserveService) GetAll(request *dto.GetAllReserveRequest) (equipments []*model.Reserve, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	equipments, err = s.reserveRepo.GetAll(ctx, &dto.GetAllReserveRequest{})
	if err != nil {
		return nil, fmt.Errorf("удаление брони: %w", err)
	}

	return equipments, err
}

func (s ReserveService) Add(request *dto.AddReserveRequest) (err error) {
	// TODO: добавить транзакцию
	//if request.EquipmentId == nil {
	//	return fmt.Errorf("слайс оборудования пуст: %w", err)
	//}

	if request.ProducerId < 0 {
		return fmt.Errorf("id продюсера меньше 0: %w", err)
	}

	if request.RoomId < 1 {
		return fmt.Errorf("id комнаты меньше 1: %w", err)
	}

	if request.InstrumentalistId < 0 {
		return fmt.Errorf("id инструменталиста меньше 0: %w", err)
	}

	if request.UserId < 1 {
		return fmt.Errorf("id пользователя меньше 1: %w", err)
	}

	for _, equipment := range request.EquipmentId {
		if equipment < 1 {
			return fmt.Errorf("id оборудования меньше 1: %w", err)
		}
	}

	//if request.TimeInterval
	if request.TimeInterval.StartTime.Unix() >= request.TimeInterval.EndTime.Unix() {
		return fmt.Errorf("время начала больше времени конца: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.reserveRepo.Add(ctx, &dto.AddReserveRequest{
		UserId:            request.UserId,
		RoomId:            request.RoomId,
		EquipmentId:       request.EquipmentId,
		ProducerId:        request.ProducerId,
		InstrumentalistId: request.InstrumentalistId,
		TimeInterval:      request.TimeInterval,
	})
	if err != nil {
		return fmt.Errorf("добавление брони: %w", err)
	}

	return err
}

func (s ReserveService) Delete(request *dto.DeleteReserveRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("id меньше 1: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.reserveRepo.Delete(ctx, &dto.DeleteReserveRequest{
		Id: request.Id,
	})
	if err != nil {
		return fmt.Errorf("удаление брони: %w", err)
	}

	return err
}
