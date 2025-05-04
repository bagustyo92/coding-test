package handler

import (
	"fmt"
	"net/http"

	"wallet/internal/service"

	"github.com/gin-gonic/gin"
)

type DisbursementHandler struct {
	service service.DisbursementService
}

func NewDisbursementHandler(s service.DisbursementService) *DisbursementHandler {
	return &DisbursementHandler{service: s}
}

type disbursementRequest struct {
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

type getUserReq struct {
	UserID string `json:"user_id"`
}

func (h *DisbursementHandler) Disburse(c *gin.Context) {
	var req disbursementRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.service.Disburse(req.UserID, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Disbursement successful",
		"disbursed_to":  user.Name,
		"amount":        req.Amount,
		"remaining_bal": user.Balance,
	})
}

func (h *DisbursementHandler) CheckBalance(c *gin.Context) {
	var req getUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user, err := h.service.GetUserByID(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       fmt.Sprintf("Hi %s, your balance is Rp %.2f", user.Name, user.Balance),
		"user_name":     user.Name,
		"remaining_bal": user.Balance,
	})
}
