package impl

import (
	"app/src/internal/model"
	"app/src/internal/model/dto"
	repositoryInterface "app/src/internal/repository/interface"
	serviceInterface "app/src/internal/service/interface"
	"context"
	"fmt"
	"slices"
)

type ValidateTimeService struct {
	roomRepo            repositoryInterface.IRoomRepository
	equipmentRepo       repositoryInterface.IEquipmentRepository
	producerRepo        repositoryInterface.IProducerRepository
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository
	reserveRepo         repositoryInterface.IReserveRepository
	//reservedEquipmentRepo repositoryInterface.IReservedEquipmentRepository
}

func NewValidateTimeService(
	roomRepo repositoryInterface.IRoomRepository,
	equipmentRepo repositoryInterface.IEquipmentRepository,
	producerRepo repositoryInterface.IProducerRepository,
	instrumentalistRepo repositoryInterface.IInstrumentalistRepository,
	reserveRepo repositoryInterface.IReserveRepository,
	// reservedEquipmentRepo repositoryInterface.IReservedEquipmentRepository,
) serviceInterface.IValidateTimeService {
	return &ValidateTimeService{
		roomRepo:            roomRepo,
		equipmentRepo:       equipmentRepo,
		producerRepo:        producerRepo,
		instrumentalistRepo: instrumentalistRepo,
		reserveRepo:         reserveRepo,
		//reservedEquipmentRepo: reservedEquipmentRepo,
	}
}

func (s ValidateTimeService) GetSuitableStuff(request *dto.GetSuitableStuffRequest) (
	notReservedRooms []*model.Room,
	notReservedEquipments [][]*model.Equipment,
	notReservedProducers []*model.Producer,
	notReservedInstrumentalists []*model.Instrumentalist,
	err error) {

	if request.StudioId < 1 {
		return nil, nil, nil, nil, fmt.Errorf("id студии меньше 1")
	}

	ctx := context.Background() //, cancel := context.WithTimeout(context.Background(), cmd.TimeOut*time.Second)
	//defer cancel()
	notReservedRooms, err = s.getNotReservedRooms(ctx, &dto.GetNotReservedRoomsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		return nil,
			nil,
			nil,
			nil,
			fmt.Errorf("поиск доступных комнат")
	}

	notReservedEquipments, err = s.getNotReservedEquipments(ctx, &dto.GetNotReservedEquipmentsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного оборудования")
	}

	notReservedProducers, err = s.getNotReservedProducers(ctx, &dto.GetNotReservedProducersRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного продюсера")
	}

	notReservedInstrumentalists, err = s.getNotReservedInstrumentalists(ctx, &dto.GetNotReservedInstrumentalistsRequest{
		ChoosenInterval: request.TimeInterval,
		StudioId:        request.StudioId,
	})
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("поиск доступного инструменталиста")
	}

	return notReservedRooms, notReservedEquipments, notReservedProducers, notReservedInstrumentalists, err
}

func (s ValidateTimeService) getNotReservedRooms(ctx context.Context, request *dto.GetNotReservedRoomsRequest) (notReservedRooms []*model.Room, err error) {

	if request.StudioId <= 0 {
		return nil, fmt.Errorf("id студии меньше 1")
	}

	//if request.StartTime.Unix() >= request.EndTime.Unix() {
	//	return nil, fmt.Errorf("")
	//}
	//request.ChoosenInterval

	rooms, err := s.roomRepo.GetByStudio(ctx, &dto.GetRoomByStudioRequest{
		StudioId: request.StudioId,
	})
	if err != nil {
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, room := range rooms {
		reserveFlag := false
		//fmt.Println("searching for producer:", producer.Id)
		//fmt.Println("id -- ", room.Id)

		if int64(request.ChoosenInterval.StartTime.Hour()) >= room.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= room.EndHour {
			reserves, err := s.reserveRepo.GetByRoomId(ctx, &dto.GetReserveByRoomIdRequest{RoomId: room.Id}) //TODO: спросить, нормально ли бегать в бд при каждой итерации
			if err != nil {
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}

			}
			if reserveFlag == false {
				notReservedRooms = append(notReservedRooms, room)
				//fmt.Println("=================", room)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	return notReservedRooms, err
}

func (s ValidateTimeService) getNotReservedProducers(ctx context.Context, request *dto.GetNotReservedProducersRequest) (notReservedProducers []*model.Producer, err error) {
	//TODO: сделать возможность обработки броней переходящих через 00:00

	if request.StudioId <= 0 {
		return nil, fmt.Errorf("id студии меньше 1")
	}

	producers, err := s.producerRepo.GetByStudio(ctx, &dto.GetProducerByStudioRequest{
		StudioId: request.StudioId,
	})
	if err != nil {
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, producer := range producers {
		reserveFlag := false

		if int64(request.ChoosenInterval.StartTime.Hour()) >= producer.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= producer.EndHour {
			reserves, err := s.reserveRepo.GetByProducerId(ctx, &dto.GetReserveByProducerIdRequest{ProducerId: producer.Id}) //TODO: спросить, нормально ли бегать в бд при каждой итерации
			if err != nil {
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}
				//reserveStartHour := int64(reserve.StartTime.Hour())
				//reserveEndHour := int64(reserve.EndTime.Hour())
				//
				////fmt.Println("	looking reserve:", reserve.Id)
				//if (choosenStartHour >= reserveStartHour && choosenStartHour < reserveEndHour) ||
				//	(choosenEndHour <= reserveEndHour && choosenEndHour > reserveStartHour) ||
				//	(choosenStartHour <= reserveStartHour && choosenEndHour >= reserveEndHour) {
				//	reserveFlag = true
				//	break
				//}
			}
			if reserveFlag == false {
				notReservedProducers = append(notReservedProducers, producer)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	//for i, _ := range notReservedProducers {
	//	fmt.Println(notReservedProducers[i])
	//}
	return notReservedProducers, err
}

func (s ValidateTimeService) getNotReservedInstrumentalists(ctx context.Context, request *dto.GetNotReservedInstrumentalistsRequest) (notReservedInstrumentalists []*model.Instrumentalist, err error) {

	if request.StudioId <= 0 {
		return nil, fmt.Errorf("id студии меньше 1")
	}

	//if request.StartTime.Unix() >= request.EndTime.Unix() {
	//	return nil, fmt.Errorf("время начала больше времени конца: %w", err)
	//}
	//
	//choosenStartHour := int64(request.StartTime.Hour())
	//choosenEndHour := int64(request.EndTime.Hour())

	instrumentalists, err := s.instrumentalistRepo.GetByStudio(ctx, &dto.GetInstrumentalistByStudioRequest{
		StudioId: request.StudioId,
	})
	if err != nil {
		return nil, fmt.Errorf("get all rooms error")
	}

	for _, instrumentalist := range instrumentalists {
		reserveFlag := false
		//fmt.Println("searching for producer:", producer.Id)

		if int64(request.ChoosenInterval.StartTime.Hour()) >= instrumentalist.StartHour &&
			int64(request.ChoosenInterval.EndTime.Hour()) <= instrumentalist.EndHour {
			reserves, err := s.reserveRepo.GetByInstrumentalistId(ctx, &dto.GetReserveByInstrumentalistIdRequest{InstrumentalistId: instrumentalist.Id}) //TODO: спросить, нормально ли бегать в бд при каждой итерации
			if err != nil {
				return nil, fmt.Errorf("get all reserves error")
			}
			for _, reserve := range reserves {
				if isIntervalsIntersect(*request.ChoosenInterval, *reserve.TimeInterval) == true {
					reserveFlag = true
					break
				}
				//reserveStartHour := int64(reserve.StartTime.Hour())
				//reserveEndHour := int64(reserve.EndTime.Hour())
				//
				////fmt.Println("	looking reserve:", reserve.Id)
				//if (choosenStartHour >= reserveStartHour && choosenStartHour < reserveEndHour) ||
				//	(choosenEndHour <= reserveEndHour && choosenEndHour > reserveStartHour) ||
				//	(choosenStartHour <= reserveStartHour && choosenEndHour >= reserveEndHour) {
				//	reserveFlag = true
				//	break
				//}
			}
			if reserveFlag == false {
				notReservedInstrumentalists = append(notReservedInstrumentalists, instrumentalist)
			}
		}
	}

	//fmt.Println(len(notReservedRooms))
	return notReservedInstrumentalists, err
}

func (s ValidateTimeService) getNotReservedEquipments(ctx context.Context, request *dto.GetNotReservedEquipmentsRequest) (notReservedEquipments [][]*model.Equipment, err error) {
	//ctx = context.Background()
	notReservedEquipments = make([][]*model.Equipment, 0)

	if request.StudioId <= 0 {
		return nil, fmt.Errorf("id студии меньше 1")
	}

	if request.ChoosenInterval.StartTime.Unix() >= request.ChoosenInterval.EndTime.Unix() {
		return nil, fmt.Errorf("")
	}

	for equipmentType := int64(model.OutOfFirstEquipment + 1); equipmentType < int64(model.OutOfLastEquipment); equipmentType++ {
		//fmt.Println(equipmentType)
		notFullTimeFreeEquipments := make([]*model.Equipment, 0)
		equipmentsAndTimes, err := s.equipmentRepo.GetNotFullTimeFreeByStudioAndType(
			ctx,
			&dto.GetEquipmentNotFullTimeFreeByStudioAndTypeRequest{
				StudioId: request.StudioId,
				Type:     equipmentType,
				TimeInterval: &model.TimeInterval{
					StartTime: request.ChoosenInterval.StartTime,
					EndTime:   request.ChoosenInterval.EndTime,
				},
			}) //select id, name, ... from equipments eq join reserved-eq r_eq with eq.id == r_eq.equipment_id where
		// Делаем запрос с between: select ... from reserved_equipment where year(date) between ...
		// Делаем крутой запрос в бд, в котором сначала ищем id в табличке бронь-оборудование, потом идем в таблицу оборужования и сравниваем id. Выбираем тех, чьи id не совпали
		// Запрос с джоином. выводим все оборудование с типом и студией из т оборудования + время брони, если оно забронированно

		for _, equipmentAndTime := range equipmentsAndTimes {
			if isIntervalsIntersect(*equipmentAndTime.TimeInterval, *request.ChoosenInterval) == false {
				notFullTimeFreeEquipments = append(notFullTimeFreeEquipments, equipmentAndTime.Equipment)
			}
		}

		if err != nil {
			//logrus.WithFields(logrus.Fields{
			//	"Internal":  "Service.validate_time",
			//	"Event":     "GetReservedByStudioAndType",
			//	"StudioId":  request.StudioId,
			//	"Type":      equipmentType,
			//	"StartTime": request.StartTime,
			//	"EndTime":   request.EndTime,
			//	"Error":     err,
			//}).Error("Ошибка получения забронированного оборудования по студии и типу из репозитория")
			return nil, fmt.Errorf("ошибка получения не аолностью забронированного оборудования по студии и типу из репозитория: %w", err)
		}
		//for i, _ := range notFullTimeFreeEquipments {
		//	fmt.Println(*notFullTimeFreeEquipments[i])
		//}
		//fmt.Println("11")

		fullTimeFreeEquipments, err := s.equipmentRepo.GetFullTimeFreeByStudioAndType(
			ctx,
			&dto.GetEquipmentFullTimeFreeByStudioAndTypeRequest{
				StudioId: request.StudioId,
				Type:     equipmentType,
			})

		//for i, _ := range fullTimeFreeEquipments {
		//	fmt.Println(*fullTimeFreeEquipments[i])
		//}
		if err != nil {
			//logrus.WithFields(logrus.Fields{
			//	"Internal":  "Service.validate_time",
			//	"Event":     "GetByStudioAndType",
			//	"StudioId":  request.StudioId,
			//	"Type":      equipmentType,
			//	"StartTime": request.StartTime,
			//	"EndTime":   request.EndTime,
			//	"Error":     err,
			//}).Error("Ошибка получения полностью свободного оборудования оборудования по студии и типу из репозитория")
			return nil, fmt.Errorf("ошибка получения полностью свободного оборудования по студии и типу из репозитория: %w", err)
		}

		notFullTimeFreeEquipments = slices.Concat(notFullTimeFreeEquipments, fullTimeFreeEquipments)
		//if len(notFullTimeFreeEquipments) != 0 {
		//	for i, _ := range notFullTimeFreeEquipments {
		//		fullTimeFreeEquipments = append(fullTimeFreeEquipments, notFullTimeFreeEquipments[i])
		//	}
		//}

		//for i, _ := range notFullTimeFreeEquipments {
		//	fmt.Println("****", notFullTimeFreeEquipments[i])
		//}

		//fullTimeFreeEquipments
		//fmt.Println("--++")
		notReservedEquipments = append(notReservedEquipments, notFullTimeFreeEquipments)
		//for i, _ := range notReservedEquipments[equipmentType-1] {
		//	notReservedEquipments[equipmentType-1] = append(notReservedEquipments[equipmentType-1], fullTimeFreeEquipments[i])
		//}
		//fmt.Println("--++")
		//notReservedEquipments[equipmentType-1] = append(notReservedEquipments[equipmentType-1], slices.Concat(fullTimeFreeEquipments, notFullTimeFreeEquipments))
		//for i, _ := range notReservedEquipments {
		//	for j, _ := range notReservedEquipments[i] {
		//		fmt.Println("----", *notReservedEquipments[i][j])
		//	}
		//}
		//fmt.Println(notReservedEquipments)
	}

	//for i, _ := range notReservedEquipments {
	//	for j, _ := range notReservedEquipments[i] {
	//		fmt.Println("=====", i, *notReservedEquipments[i][j])
	//	}
	//}

	return notReservedEquipments, err
}
