package handlers

import (
	"github.com/Flexours-Team/API/pkg/model"
	"github.com/Flexours-Team/models"
	"github.com/gin-gonic/gin"
	uuid "github.com/go-openapi/strfmt"
)

func (h *Handler) GetAbsenses(ctx *gin.Context) {
	absenses, err := h.App.DB.GetAbsenses()
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2060,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, model.ListAbsenseResponse{
		Code:    0,
		Message: "success",
		Data:    absenses,
	})

}
func (h *Handler) GetAbsenseByID(ctx *gin.Context) {
	absenseId := uuid.UUID4(ctx.Param("id"))
	absense, err := h.App.DB.GetAbsenseByID(absenseId)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2060,
			Message: err.Error(),
		})
		return
	}
	ctx.JSON(200, model.ByIDAbsenseResponse{
		Code:    0,
		Message: "success",
		Data:    absense,
	})
}
func (h *Handler) CreateAbsense(ctx *gin.Context) {
	var createAbsenseReq model.CreateAbsenseRequest
	if err := ctx.ShouldBindJSON(&createAbsenseReq); err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2050,
			Message: err.Error(),
		})
		return
	}
	//start, err :=time.Parse("2006-01-03 22:04:05", createAbsenseReq.StartDate)
	//if err != nil {
	//	ctx.JSON(400, model.BaseResponse{
	//		Code:    2050,
	//		Message: err.Error(),
	//	})
	//}
	//finish, err := time.Parse("2006-01-03 22:04:05", createAbsenseReq.EndDate)
	//if err != nil {
	//	ctx.JSON(400, model.BaseResponse{
	//		Code:    2050,
	//		Message: err.Error(),
	//	})
	//}
	absense := &models.Absense{
		StartDate    :createAbsenseReq.StartDate,
		EndDate      :createAbsenseReq.EndDate,
		AbsenseReason:createAbsenseReq.AbsenseReason,
	}
	err := h.App.DB.CreateAbsense(absense)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2052,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(200, model.BaseResponse{
		Message: "Success",
		Data: absense,
	})

}
func (h *Handler) UpdateAbsenseByID(ctx *gin.Context) {
	var updateAbsenseReq model.UpdateAbsenseRequest
	if err := ctx.ShouldBindJSON(&updateAbsenseReq); err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2050,
			Message: err.Error(),
		})
		return
	}
	absenseID := uuid.UUID4(ctx.Param("id"))
	//start, err :=time.Parse("2006-01-03 22:04:05", updateAbsenseReq.StartDate)
	//if err != nil {
	//	ctx.JSON(400, model.BaseResponse{
	//		Code:    2050,
	//		Message: err.Error(),
	//	})
	//}
	//finish, err := time.Parse("2006-01-03 22:04:05", updateAbsenseReq.EndDate)
	//if err != nil {
	//	ctx.JSON(400, model.BaseResponse{
	//		Code:    2050,
	//		Message: err.Error(),
	//	})
	//}

	absense := &models.Absense{
		StartDate    : updateAbsenseReq.StartDate,
		EndDate      : updateAbsenseReq.EndDate,
		AbsenseReason:updateAbsenseReq.AbsenseReason,
	}
	absense.ID = absenseID
	err := h.App.DB.UpdateAbsenseByID(absense)
	if err != nil {
		ctx.JSON(400, model.BaseResponse{
			Code:    2052,
			Message: err.Error(),
		})
		return
	}

	scheduleUpdated, err := h.App.DB.GetAbsenseByID(absenseID)
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
//func (h *Handler) DeleteAbsense(ctx *gin.Context) {
//
//}