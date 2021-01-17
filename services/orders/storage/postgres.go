package storage

import (
	"context"

	"delivery/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type OrderStorage struct {
	db *sqlx.DB
}

func NewOrderStorage(db *sqlx.DB) *OrderStorage {
	return &OrderStorage{
		db: db,
	}
}

func (s *OrderStorage) CreateOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	query := `INSERT INTO orders(ID, number, date) VALUES ($1, $2, $3)`

	o.ID = uuid.New().String()
	o.Manager = m

	if _, err := s.db.ExecContext(ctx, query, o.ID, o.Number, o.Date); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *OrderStorage) GetOrders(ctx context.Context, m *models.Manager) ([]*models.Order, error) {

	var order models.Order
	var list []*models.Order

	rows, err := s.db.Queryx("SELECT * FROM orders WHERE manager_id = $1 ORDER BY date DESC;", m.ID)
	for rows.Next() {
		if err = rows.StructScan(&order); err != nil {
			return nil, errors.WithStack(err)
		}
		list = append(list, &order)
	}
	return list, nil
}

func (s *OrderStorage) DeleteOrder(ctx context.Context, m *models.Manager, o *models.Order) error {
	query := `DELETE FROM orders WHERE id = $1;`

	if _, err := s.db.ExecContext(ctx, query, o.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
