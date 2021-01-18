package storage

import (
	"context"
	"delivery/models"
	"github.com/stretchr/testify/mock"
)

type OrderStorageMock struct {
	mock.Mock
}

func (s *OrderStorageMock) Insert(ctx context.Context, m *models.User, o *models.Order) error {
	args := s.Called(m, o)
	return args.Error(0)
}

func (s *OrderStorageMock) List(ctx context.Context, m *models.User) ([]*models.Order, error) {
	args := s.Called(m)
	return args.Get(0).([]*models.Order), args.Error(1)
}

func (s *OrderStorageMock) Delete(ctx context.Context, m *models.User, o *models.Order) error {
	args := s.Called(m, o)
	return args.Error(0)
}