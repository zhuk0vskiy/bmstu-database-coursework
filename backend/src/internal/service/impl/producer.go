package impl

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
	repositoryInterface "app/src/internal/repository/interface"
	serviceInterface "app/src/internal/service/interface"
	"context"
	"fmt"
)

type ProducerService struct {
	producerRepo repositoryInterface.IProducerRepository
	reserveRepo  repositoryInterface.IReserveRepository
}

func NewProducerService(
	producerRepo repositoryInterface.IProducerRepository,
	reserveRepo repositoryInterface.IReserveRepository) serviceInterface.IProducerService {
	return &ProducerService{
		producerRepo: producerRepo,
		reserveRepo:  reserveRepo,
	}
}

func (s ProducerService) GetByStudio(request *dto.GetProducerByStudioRequest) (producers []*model.Producer, err error) {
	if request.StudioId < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() // , cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	producers, err = s.producerRepo.GetByStudio(ctx, &dto.GetProducerByStudioRequest{
		StudioId: request.StudioId,
	})

	if err != nil {
		return nil, fmt.Errorf("получение продюсеров по студии: %w", err)
	}
	return producers, err
}

func (s ProducerService) Get(request *dto.GetProducerRequest) (producer *model.Producer, err error) {
	if request.Id < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	producer, err = s.producerRepo.Get(ctx, &dto.GetProducerRequest{
		Id: request.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("получение продюсера по id: %w", err)
	}

	return producer, err
}

//func (s ProducerService) GetAll() (producers []*domain.Producer, err error) {
//	producers, err = s.producerRepo.GetAll()
//
//	if err != nil {
//		return nil, fmt.Errorf("получение всех продюсеров: %w", err)
//	}
//
//	return producers, nil
//}

func (s ProducerService) Update(request *dto.UpdateProducerRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	isReserve, err := s.reserveRepo.IsProducerReserve(ctx, &dto.IsProducerReserveRequest{
		ProducerId: request.Id,
	})
	if err != nil {
		return fmt.Errorf("получение всех броней: %w", err)
	}

	if isReserve == true {
		return fmt.Errorf("нельзя обновить продюсера, тк на него есть бронь: %w", err)
	}

	if request.Name == "" {
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 0 {
		return fmt.Errorf("id студии меньше 0")
	}
	//
	//if request.StartTime.s

	if request.StartHour == request.EndHour {
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.EndHour < 0 || request.StartHour < 0 {
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	err = s.producerRepo.Update(ctx, &dto.UpdateProducerRequest{
		Id:        request.Id,
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		return fmt.Errorf("обновление продюсера: %w", err)
	}

	return err
}

func (s ProducerService) Add(request *dto.AddProducerRequest) (err error) {

	if request.Name == "" {
		return fmt.Errorf("пустое имя: %w", err)
	}

	if request.StudioId < 1 {
		return fmt.Errorf("id студии меньше 0")
	}

	if request.StartHour == request.EndHour {
		return fmt.Errorf("время начала работы равно времени конца")
	}

	if request.EndHour > 23 || request.StartHour > 23 ||
		request.StartHour < 0 || request.EndHour < 0 {
		return fmt.Errorf("время конца/начала работы не входит в размер суток (от 00 до 23)")
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	err = s.producerRepo.Add(ctx, &dto.AddProducerRequest{
		Name:      request.Name,
		StudioId:  request.StudioId,
		StartHour: request.StartHour,
		EndHour:   request.EndHour,
	})
	if err != nil {
		return fmt.Errorf("добавление продюсера: %w", err)
	}

	return err
}

func (s ProducerService) Delete(request *dto.DeleteProducerRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	err = s.producerRepo.Delete(ctx, &dto.DeleteProducerRequest{
		Id: request.Id,
	})

	if err != nil {
		return fmt.Errorf("удаление продюсера: %w", err)
	}

	return err
}
