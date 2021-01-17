package orders

import (
	"context"
	"delivery/models"
)

type UseCase interface {
	CreateOrder(ctx context.Context, m *models.Manager, o *CreateOrder) error
	GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error)
	DeleteOrder(ctx context.Context, m *models.Manager, o *DeleteOrder) error
}
