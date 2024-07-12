package dto

import (
	"github.com/prometheus/client_golang/prometheus"
)

type IncUsersOnlineRequest struct {
	usersOnline prometheus.Gauge
}
