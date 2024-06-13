package dto

import "app/src/internal/model"

type NewTimeIntervalRequest struct {
	StartTime *model.Time
	EndTime   *model.Time
}

type NewTimeRequest struct {
	Time *model.Time
}
