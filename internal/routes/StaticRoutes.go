package routes

import (
	"Calculator/internal/handler"

	"github.com/gin-gonic/gin"
)

func StaticRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/cal", handler.CalculateHandler())
}
