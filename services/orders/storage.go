package orders

import (
	"context"
	"delivery/models"
)

type Storage interface {
	// Insert new order to storage
	Insert(context.Context, *models.User,  *models.Order) error
	// Delete existing order from storage
	Delete(context.Context, *models.User,  *models.Order) error
	// List of orders for user from storage
	List(context.Context, *models.User) ([]*models.Order, error)
}
