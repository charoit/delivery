package storage

import (
	"context"
	"delivery/models"
	"github.com/stretchr/testify/mock"
)

type OrderStorageMock struct {
	mock.Mock
}

func (s *OrderStorageMock) CreateOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	args := s.Called(m, o)
	return args.Error(0)
}

func (s *OrderStorageMock) GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error) {
	args := s.Called(m)
	return args.Get(0).([]*models.Order), args.Error(1)
}

func (s *OrderStorageMock) DeleteOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	args := s.Called(m, o)
	return args.Error(0)
}