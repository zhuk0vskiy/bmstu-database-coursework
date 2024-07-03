func IncUsersOnline(request *dto.IncUsersOnlineRequest) (err error) {
	request.usersOnline.Inc()
}