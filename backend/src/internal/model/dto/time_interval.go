package dto

import "github.com/zhuk0vskiy/bmstu-database-coursework/backend/src/internal/model"

type NewTimeIntervalRequest struct {
	StartTime *model.Time
	EndTime   *model.Time
}

type NewTimeRequest struct {
	Time *model.Time
}
