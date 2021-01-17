package storage

import (
	"context"
	"delivery/models"
	"github.com/pkg/errors"
	"sync"
)

var ErrOrderNotFound = errors.New("order not found")

type OrderLocalStorage struct {
	*sync.RWMutex
	orders map[string]*models.Order
}

func NewBookmarkLocalStorage() *OrderLocalStorage {
	return &OrderLocalStorage{
		RWMutex: &sync.RWMutex{},
		orders: make(map[string]*models.Order),
	}
}

func (s *OrderLocalStorage) CreateOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	o.Manager = m

	s.Lock()
	s.orders[o.ID] = o
	s.Unlock()

	return nil
}

func (s *OrderLocalStorage) GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)

	s.Lock()
	for _, o := range s.orders {
		if o.Manager == m {
			orders = append(orders, o)
		}
	}
	s.Unlock()

	return orders, nil
}

func (s *OrderLocalStorage) DeleteOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	s.Lock()
	defer s.Unlock()

	order, ex := s.orders[o.ID]
	if ex && order.Manager == m {
		delete(s.orders, o.ID)
		return nil
	}

	return ErrOrderNotFound
}

