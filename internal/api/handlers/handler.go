package handlers

import (
	"github.com/Flexours-Team/API/pkg/application"
	"github.com/gin-gonic/gin"
)

// Handler model
type Handler struct {
	App     application.Application
}

// Get - Handler initializer
func Get(app application.Application) *Handler {
	var handler Handler
	handler.App = app

	return &Handler{
		App:     app,
	}
}

func (h *Handler) CORSMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, grant_type, client_id, client_secret, channel_secret, Grant-Type, Client-ID, Client-Secret, Channel-Secret, merchant_id, Merchant-ID")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(204)
		return
	}
	ctx.Next()
}

