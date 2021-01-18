package api

import (
	"delivery/auth"
	"net/http"

	"delivery/models"
	"delivery/services/orders"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service orders.Service
}

func NewHandler(svc orders.Service) *handler {
	return &handler{
		service: svc,
	}
}

func (h *handler) OrderCreate(c *gin.Context) {
	order := &orders.Order{}
	if err := c.BindJSON(order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.service.Create(c.Request.Context(), user, order); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *handler) OrderList(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*models.User)

	list, err := h.service.List(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *handler) OrderRemove(c *gin.Context) {
	inp := &orders.Remove{}
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.User)

	if err := h.service.Remove(c.Request.Context(), user, inp); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}
