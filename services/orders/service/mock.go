package usecase

import (
	"context"
	"delivery/models"
	"delivery/services/orders"
	"github.com/stretchr/testify/mock"
)

type OrderUseCaseMock struct {
	mock.Mock
}

func (uc *OrderUseCaseMock) CreateOrder(ctx context.Context, m *models.User, o *orders.CreateOrder) error {
	args := uc.Called(m, o)
	return args.Error(0)
}

func (uc *OrderUseCaseMock) GetOrders(ctx context.Context, m *models.User) ([]*models.Order, error) {
	args := uc.Called(m)
	return args.Get(0).([]*models.Order), args.Error(1)
}

func (uc *OrderUseCaseMock) DeleteOrder(ctx context.Context, m *models.User, o *orders.DeleteOrder) error {
	args := uc.Called(m, o)
	return args.Error(0)
}

