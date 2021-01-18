package service

import (
	"context"
	"delivery/models"
	"delivery/services/orders"
	"github.com/stretchr/testify/mock"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Create(ctx context.Context, user *models.User, order *orders.Order) error {
	args := s.Called(user, order)
	return args.Error(0)
}

func (s *ServiceMock) Remove(ctx context.Context, user *models.User, order *orders.Remove) error {
	args := s.Called(user, order)
	return args.Error(0)
}

func (s *ServiceMock) List(ctx context.Context, user *models.User) (*orders.List, error) {
	args := s.Called(user)
	list := args.Get(0).([]*models.Order)
	return s.ToOrderList(list), args.Error(1)
}

func (s *ServiceMock) ToOrderList(src []*models.Order) *orders.List {
	dst := &orders.List{
		Orders: make([]*orders.Order, len(src)),
	}

	for i, o := range src {
		dst.Orders[i] = &orders.Order{
			ID:          o.ID,
			Number:      o.Number,
			DeliveryId:  o.Delivery.ID,
			RecipientId: o.Recipient.ID,
		}
	}
	return dst
}
