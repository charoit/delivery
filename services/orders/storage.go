package orders

import (
	"context"
	"delivery/models"
)

type Storage interface {
	CreateOrder(ctx context.Context, m *models.Manager, o *models.Order) error
	GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error)
	DeleteOrder(ctx context.Context, m *models.Manager, o *models.Order) error
}
