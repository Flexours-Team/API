package model

import (
	"github.com/Flexours-Team/models"
	"time"
)

type ByIDScheduleResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    models.Schedule `json:"data"`
}

type ListScheduleResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []models.Schedule `json:"data"`
}

type StartTimeRequest struct {
	StartDate     	*time.Time `json:"startDate"`
	RegType 		models.RegistrationTypes `json:"regType" binding:"required"`
}

type UpdateScheduleRequest struct {
	EndDate       	string `json:"endDate" gorm:"default:now()"`
}

type GetDayHoursResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    float64 `json:"data"`
}
