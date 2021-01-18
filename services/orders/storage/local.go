package storage

import (
	"context"
	"delivery/models"
	"github.com/pkg/errors"
	"strings"
	"sync"
)

var ErrOrderNotFound = errors.New("order not found")

type localStorage struct {
	*sync.RWMutex
	orders map[string]*models.Order
}

func NewLocalStorage() *localStorage {
	return &localStorage{
		RWMutex: &sync.RWMutex{},
		orders: make(map[string]*models.Order),
	}
}

func (s *localStorage) Insert(ctx context.Context, user *models.User, order *models.Order) error {
	order.Manager = user

	s.Lock()
	s.orders[order.ID] = order
	s.Unlock()

	return nil
}

func (s *localStorage) List(ctx context.Context, user *models.User) ([]*models.Order, error) {
	orders := make([]*models.Order, 0)

	s.RLock()
	for _, o := range s.orders {
		if strings.EqualFold(o.Manager.ID, user.ID) {
			orders = append(orders, o)
		}
	}
	s.RUnlock()

	return orders, nil
}

func (s *localStorage) Delete(ctx context.Context, user *models.User, order *models.Order) error {
	s.Lock()
	defer s.Unlock()

	o, found := s.orders[order.ID]
	if found && strings.EqualFold(o.Manager.ID, user.ID) {
		delete(s.orders, o.ID)
		return nil
	}
	return errors.WithStack(ErrOrderNotFound)
}

