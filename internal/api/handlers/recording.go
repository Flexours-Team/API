package handlers

import (
	"fmt"
	"github.com/Flexours-Team/API/pkg/model"
	"github.com/Flexours-Team/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/go-openapi/strfmt"
	"time"
)

func (h *Handler) StartTime(ctx *gin.Context) {
	var createScheduleReq model.StartTimeRequest
	if err := ctx.ShouldBindJSON(&createScheduleReq); err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2050,
			Message: err.Error(),
		})
		return
	}
	schedule := &models.Schedule{
		RegType: createScheduleReq.RegType,
	}
	err := h.App.DB.CreateStartScheduling(schedule)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2052,
			Message: err.Error(),
		})
		return
	}
	fmt.Println("HEREE: ", schedule)
	ctx.JSON(200, model.BaseResponse{
		Message: "Success",
		Data: schedule,
	})
}

func (h *Handler) EndTime(ctx *gin.Context) {
	//var updateScheduleReq model.UpdateScheduleRequest
	recordID := uuid.UUID4(ctx.Param("id"))
	schedule := &models.Schedule{
		EndDate: time.Now(),
	}
	schedule.ID = recordID
	err := h.App.DB.CreateEndScheduling(schedule)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2052,
			Message: err.Error(),
		})
		return
	}

	scheduleUpdated, err := h.App.DB.GetSchedulingByID(recordID)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2052,
			Message: err.Error(),
		})
		return
	}


	ctx.JSON(200, model.BaseResponse{
		Message: "Success",
		Data: scheduleUpdated,
	})
}

func (h *Handler) GetSchedulings(ctx *gin.Context) {
	schedules, err := h.App.DB.GetSchedulings()
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2060,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, model.ListScheduleResponse{
		Code:    0,
		Message: "success",
		Data:    schedules,
	})
}
func (h *Handler) GetScheduleByID(ctx *gin.Context) {
	scheduleId := uuid.UUID4(ctx.Param("id"))
	schedule, err := h.App.DB.GetSchedulingByID(scheduleId)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2060,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, model.ByIDScheduleResponse{
		Code:    0,
		Message: "success",
		Data:    schedule,
	})
}

func (h *Handler) GetDayHours(ctx *gin.Context) {
	t := time.Now()
	totalHours, err := h.App.DB.GetDayHours(fmt.Sprintf("%02d", t.Day()), fmt.Sprintf("%02d", t.Day()+1) )
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2060,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, model.GetDayHoursResponse{
		Code:    0,
		Message: "success",
		Data:    totalHours,
	})
}