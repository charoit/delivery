package api

import (
	"delivery/services/orders"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc orders.UseCase) {
	h := NewHandler(uc)
	orders := router.Group("/orders")
	{
		orders.POST("", h.Create)
		orders.GET("", h.Get)
		orders.DELETE("", h.Delete)
	}
}

