package routes

import (
	"wallet/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handler.DisbursementHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/disburse", h.Disburse)
	router.POST("/check_balance", h.CheckBalance)
	return router
}
