package storage

import (
	"context"
	"delivery/models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrder(t *testing.T) {
	id := "id"
	user := &models.Manager{ID: id}

	s := NewBookmarkLocalStorage()

	for i := 0; i < 10; i++ {
		o := &models.Order{
			ID:     fmt.Sprintf("id%d", i),
			Manager: user,
		}

		err := s.CreateOrder(context.Background(), user, o)
		assert.NoError(t, err)
	}

	returnedBookmarks, err := s.GetOrders(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedBookmarks))
}

func TestDeleteOrder(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.Manager{ID: id1}
	user2 := &models.Manager{ID: id2}

	o := &models.Order{ID: "oID", Manager: user1}

	s := NewBookmarkLocalStorage()

	err := s.CreateOrder(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.DeleteOrder(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.CreateOrder(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.DeleteOrder(context.Background(), user2, o)
	assert.Error(t, err)
	assert.Equal(t, err, ErrOrderNotFound)
}

