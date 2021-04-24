package api

import (
	"github.com/Flexours-Team/API/internal/api/handlers"
	"github.com/Flexours-Team/API/pkg/application"
	"github.com/gin-gonic/gin"
)

func New(app application.Application) (*gin.Engine, error) {
	router := gin.Default()
	handler := handlers.Get(app)

	router.Use(handler.CORSMiddleware)
	v1 := router.Group("/v1")

	v1.POST("/start", handler.StartTime)
	v1.POST("/end/:id", handler.EndTime)
	v1.GET("/schedules", handler.GetSchedulings)
	v1.GET("/schedules/:id", handler.GetScheduleByID)

	v1.GET("/absense", handler.GetAbsenses)
	v1.GET("/absense/:id", handler.GetAbsenseByID)
	v1.POST("/absense", handler.CreateAbsense)
	v1.PUT("/absense/:id", handler.UpdateAbsenseByID)

	v1.GET("/get-day-hours", handler.GetDayHours)

	return router, nil
}
