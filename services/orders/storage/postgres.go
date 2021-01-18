package storage

import (
	"context"

	"delivery/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) Insert(ctx context.Context, user *models.User, order *models.Order) error {
	query := `INSERT INTO orders(ID, number, date) VALUES ($1, $2, $3)`

	order.ID = uuid.New().String()
	order.Manager = user

	if _, err := s.db.ExecContext(ctx, query, order.ID, order.Number, order.Date); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s *storage) List(ctx context.Context, user *models.User) ([]*models.Order, error) {

	var order models.Order
	var list []*models.Order

	rows, err := s.db.Queryx("SELECT * FROM orders WHERE manager_id = $1 ORDER BY date DESC;", user.ID)
	for rows.Next() {
		if err = rows.StructScan(&order); err != nil {
			return nil, errors.WithStack(err)
		}
		list = append(list, &order)
	}
	return list, nil
}

func (s *storage) Delete(ctx context.Context, user *models.User, order *models.Order) error {
	query := `DELETE FROM orders WHERE id = $1;`

	if _, err := s.db.ExecContext(ctx, query, order.ID); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
