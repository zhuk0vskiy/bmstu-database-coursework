package impl

import (
	"fmt"
	"monitoring/prometheus/dto"
)

func IncUsersOnline(request *dto.IncUsersOnlineRequest) (err error) {
	
	if err != nil {
		return fmt.Errorf("нелья увеличить количество пользователей онлайн")
	}
	return err
}
