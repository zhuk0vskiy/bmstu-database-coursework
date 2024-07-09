import ("prometheus")

type IncUsersOnlineRequest struct {
	usersOnline prometheus.NewGauge
}