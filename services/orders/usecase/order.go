package usecase

import (
	"context"
	"time"

	"delivery/models"
	"delivery/services/orders"
	"github.com/google/uuid"
)

type OrderUseCase struct {
	storage orders.Storage
}

func NewOrderUseCase(storage orders.Storage) *OrderUseCase {
	return &OrderUseCase{
		storage: storage,
	}
}

func (uc OrderUseCase) CreateOrder(ctx context.Context, m *models.Manager, o *orders.CreateOrder) error {
	order := &models.Order{
		ID:        uuid.New().String(),
		Number:    o.Number,
		Manager:   m,
		Date:      time.Now(),
		Delivery:  models.Address{},
		Recipient: models.Recipient{},
	}
	return uc.storage.CreateOrder(ctx, m, order)
}

func (uc OrderUseCase) GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error) {
	return uc.storage.GetOrders(ctx, m)
}

func (uc OrderUseCase) DeleteOrder(ctx context.Context, m *models.Manager, o *orders.DeleteOrder) error {
	order := &models.Order{
		ID: o.ID,
	}
	return uc.storage.DeleteOrder(ctx, m, order)
}
