// internal/adapters/gin/handler.go
package ginadapter

import (
	"net/http"
	"order-service/internal/app"
	"order-service/internal/domain"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	Service *app.OrderService
}

func NewOrderHandler(router *gin.Engine, svc *app.OrderService) {
	h := &OrderHandler{Service: svc}

	router.POST("/orders", h.CreateOrder)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order created"})
}
