package orders

import (
	"context"
	"delivery/models"
)

type Order struct {
	ID          string `json:"id"`
	Number      string `json:"number"`
	DeliveryId  string `json:"delivery_id"`
	RecipientId string `json:"recipient_id"`
}

type Remove struct {
	ID string `json:"id"`
}

type List struct {
	Orders []*Order `json:"orders"`
}

type Service interface {
	// Create new order from API
	Create(context.Context, *models.User, *Order) error
	// Remove order from API
	Remove(context.Context, *models.User, *Remove) error
	// List orders from API
	List(context.Context, *models.User) (*List, error)
	// ToOrderList convert model order list to API order list
	ToOrderList([]*models.Order) *List
}
