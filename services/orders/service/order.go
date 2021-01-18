package usecase

import (
	"context"
	"github.com/pkg/errors"
	"time"

	"delivery/models"
	"delivery/services/orders"
	"github.com/google/uuid"
)

type service struct {
	storage orders.Storage
}

// NewService create API interface
func NewService(storage orders.Storage) *service {
	return &service{
		storage: storage,
	}
}

// Create new order from API
func (s *service) Create(ctx context.Context, user *models.User, order *orders.Order) error {
	o := &models.Order{
		ID:        uuid.New().String(),
		Number:    order.Number,
		Manager:   user,
		Date:      time.Now(),
		Delivery:  models.Address{},
		Recipient: models.Recipient{},
	}
	return s.storage.Insert(ctx, user, o)
}

// Remove order from API
func (s *service) Remove(ctx context.Context, user *models.User, order *orders.Remove) error {
	o := &models.Order{
		ID: order.ID,
	}
	return s.storage.Delete(ctx, user, o)
}

// List orders from API
func (s *service) List(ctx context.Context, user *models.User) (*orders.List, error) {
	src, err := s.storage.List(ctx, user)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	dst := s.toOrderList(src)
	return dst, nil
}

func (s *service) toOrderList(src []*models.Order) *orders.List {
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