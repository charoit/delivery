package api

import (
	"delivery/services/orders"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, svc orders.Service) {
	h := NewHandler(svc)
	order := router.Group("/orders")
	{
		order.POST("", h.OrderCreate)
		order.DELETE("", h.OrderRemove)
		order.GET("", h.OrderList)
	}
}

