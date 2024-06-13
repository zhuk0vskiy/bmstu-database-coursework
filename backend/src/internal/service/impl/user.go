package impl

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
	repositoryInterface "app/src/internal/repository/interface"
	serviceInterface "app/src/internal/service/interface"
	"app/src/pkg/base"
	"context"
	"fmt"
)

type UserService struct {
	userRepo    repositoryInterface.IUserRepository
	reserveRepo repositoryInterface.IReserveRepository
	crypto      base.IHashCrypto
}

func NewUserService(
	userRepo repositoryInterface.IUserRepository,
	reserveRepo repositoryInterface.IReserveRepository,
	crypto base.IHashCrypto,
) serviceInterface.IUserService {
	return &UserService{
		userRepo:    userRepo,
		reserveRepo: reserveRepo,
		crypto:      crypto,
	}
}

func (s UserService) GetReserves(request *dto.GetUserReservesRequest) (reserves []*model.Reserve, err error) {
	if request.Id < 1 {
		return nil, fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()

	reserves, err = s.reserveRepo.GetUserReserves(ctx, &dto.GetUserReservesRequest{
		Id: request.Id,
	})

	if err != nil {
		return nil, fmt.Errorf("получение всех броней: %w", err)
	}

	return reserves, nil
}

func (s UserService) Get(request *dto.GetUserRequest) (user *model.User, err error) {
	if request.Id < 1 {
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	user, err = s.userRepo.Get(ctx, &dto.GetUserRequest{
		Id: request.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по id: %w", err)
	}

	return user, err
}

func (s UserService) GetByLogin(request *dto.GetUserByLoginRequest) (user *model.User, err error) {
	if request.Login == "" {
		return nil, fmt.Errorf("получение пользователя по логину: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	user, err = s.userRepo.GetByLogin(ctx, &dto.GetUserByLoginRequest{
		Login: request.Login,
	})
	if err != nil {
		return nil, fmt.Errorf("получение пользователя по логину: %w", err)
	}

	return user, err
}

//func (s UserService) GetByLogin(login string) (user *domain.User, err error) {
//	user, err = s.userRepo.GetByLogin(login)
//	if err != nil {
//		return nil, fmt.Errorf("получение пользователя по login: %w", err)
//	}
//
//	return user, err
//}

func (s UserService) Update(request *dto.UpdateUserRequest) (err error) {
	//reserves, err := s.reserveRepo.GetAll()
	//if err != nil {
	//	return fmt.Errorf("получение всех броней: %w", err)
	//}
	//
	//for _, reserve := range reserves {
	//	if reserve.int == user.Id {
	//		return fmt.Errorf("нельзя обновить пользователя, тк на нее есть бронь")
	//	}
	//}
	if request.Id < 1 {
		return fmt.Errorf("id < 1: %w", err)
	}

	if request.Login == "" {
		return fmt.Errorf("пустой логин: %w", err)
	}

	if request.Password == "" {
		return fmt.Errorf("пустой пароль: %w", err)
	}

	if request.FirstName == "" {
		return fmt.Errorf("пустое имя %w", err)
	}

	if request.SecondName == "" {
		return fmt.Errorf("пустая фамилия %w", err)
	}

	if request.ThirdName == "" {
		return fmt.Errorf("пустое отчество %w", err)
	}

	hashedPassword, err := s.crypto.GenerateHashPass(request.Password)

	if err != nil {
		return fmt.Errorf("генерация хэша: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.userRepo.Update(ctx, &dto.UpdateUserRequest{
		Id:         request.Id,
		Login:      request.Login,
		Password:   hashedPassword,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
		ThirdName:  request.ThirdName,
	})
	if err != nil {
		return fmt.Errorf("обновление пользователя %w", err)
	}

	return err
}

func (s UserService) Delete(request *dto.DeleteUserRequest) (err error) {
	if request.Id < 1 {
		return fmt.Errorf("неверный id: %w", err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	ctx := context.Background()
	err = s.userRepo.Delete(ctx, &dto.DeleteUserRequest{
		Id: request.Id,
	})
	if err != nil {
		return fmt.Errorf("удаление пользователя: %w", err)
	}

	return err
}
