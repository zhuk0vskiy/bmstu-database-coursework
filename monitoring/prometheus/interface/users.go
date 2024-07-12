package _interface

import "monitoring/prometheus/dto"

type IMonitorService interface {
	IncUsersOnline(*dto.IncUsersOnlineRequest) error
}
