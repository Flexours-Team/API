package model

import (
	"github.com/Flexours-Team/models"
	"time"
)

type ByIDAbsenseResponse struct {
	Code    int            `json:"code"`
	Message string         `json:"message"`
	Data    models.Absense `json:"data"`
}

type ListAbsenseResponse struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    []models.Absense `json:"data"`
}

type CreateAbsenseRequest struct {
	StartDate     	*time.Time `json:"startDate" binding:"required"`
	EndDate *time.Time `json:"endDate" binding:"required"`
	AbsenseReason string `json:"absenseReason" binding:"required"`
}

type UpdateAbsenseRequest struct {
	StartDate     	*time.Time `json:"startDate"`
	EndDate *time.Time `json:"endDate"`
	AbsenseReason string `json:"absenseReason"`
}
