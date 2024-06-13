package impl

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
	repositoryInterface "app/src/internal/repository/interface"
	serviceInterface "app/src/internal/service/interface"
	"context"
	"fmt"
)

type StudioService struct {
	studioRepo repositoryInterface.IStudioRepository
}

func NewStudioService(studioRepo repositoryInterface.IStudioRepository) serviceInterface.IStudioService {
	return &StudioService{
		studioRepo: studioRepo,
	}
}

func (s StudioService) Update(request *dto.UpdateStudioRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неправильный id: %w", err)
	}

	if request.Name == "" {
		return fmt.Errorf("пустое название: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.studioRepo.Update(ctx, &dto.UpdateStudioRequest{
		Id:   request.Id,
		Name: request.Name,
	})
	if err != nil {
		return fmt.Errorf("обновление студии: %w", err)
	}

	return err
}

func (s StudioService) Get(request *dto.GetStudioRequest) (studio *model.Studio, err error) {
	if request.Id < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	studio, err = s.studioRepo.Get(ctx, &dto.GetStudioRequest{
		Id: request.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("получение студии по id: %w", err)
	}

	return studio, err
}

func (s StudioService) GetAll(request *dto.GetStudioAllRequest) (studios []*model.Studio, err error) {

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	studios, err = s.studioRepo.GetAll(ctx, &dto.GetStudioAllRequest{})

	if err != nil {
		return nil, fmt.Errorf("получение студий по id: %w", err)
	}

	return studios, err
}

func (s StudioService) Add(request *dto.AddStudioRequest) (err error) {

	if request.Name == "" {
		return fmt.Errorf("пустое имя: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.studioRepo.Add(ctx, &dto.AddStudioRequest{
		Name: request.Name,
	})

	if err != nil {
		return fmt.Errorf("добавление студии: %w", err)
	}

	return err
}

func (s StudioService) Delete(request *dto.DeleteStudioRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.studioRepo.Delete(ctx, &dto.DeleteStudioRequest{
		Id: request.Id,
	})

	if err != nil {
		return fmt.Errorf("удаление студии по id: %w", err)
	}

	return err
}
