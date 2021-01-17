package api

import (
	"delivery/auth"
	"net/http"

	"delivery/models"
	"delivery/services/orders"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase orders.UseCase
}

func NewHandler(useCase orders.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h *Handler) Create(c *gin.Context) {
	order := &orders.CreateOrder{}
	if err := c.BindJSON(order); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.Manager)

	if err := h.useCase.CreateOrder(c.Request.Context(), user, order); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) Get(c *gin.Context) {

	user := c.MustGet(auth.CtxUserKey).(*models.Manager)

	bms, err := h.useCase.GetOrders(c.Request.Context(), user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &orders.GetOrders{
		Orders: toBookmarks(bms),
	})
}

func (h *Handler) Delete(c *gin.Context) {
	inp := &orders.DeleteOrder{}
	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := c.MustGet(auth.CtxUserKey).(*models.Manager)

	if err := h.useCase.DeleteOrder(c.Request.Context(), user, inp); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func toBookmarks(bs []*models.Order) []*orders.CreateOrder {
	out := make([]*orders.CreateOrder, len(bs))

	for i, b := range bs {
		out[i] = toBookmark(b)
	}

	return out
}

func toBookmark(o *models.Order) *orders.CreateOrder {
	return &orders.CreateOrder{
		ID:          o.ID,
		Number:      o.Number,
		DeliveryId:  o.Delivery.ID,
		RecipientId: o.Recipient.ID,
	}
}
