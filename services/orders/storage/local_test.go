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
	user := &models.User{ID: id}

	s := NewLocalStorage()

	for i := 0; i < 10; i++ {
		o := &models.Order{
			ID:     fmt.Sprintf("id%d", i),
			Manager: user,
		}

		err := s.Insert(context.Background(), user, o)
		assert.NoError(t, err)
	}

	returnedBookmarks, err := s.List(context.Background(), user)
	assert.NoError(t, err)

	assert.Equal(t, 10, len(returnedBookmarks))
}

func TestDeleteOrder(t *testing.T) {
	id1 := "id1"
	id2 := "id2"

	user1 := &models.User{ID: id1}
	user2 := &models.User{ID: id2}

	o := &models.Order{ID: "oID", Manager: user1}

	s := NewLocalStorage()

	err := s.Insert(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.Delete(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.Insert(context.Background(), user1, o)
	assert.NoError(t, err)

	err = s.Delete(context.Background(), user2, o)
	assert.Error(t, err)
	assert.Equal(t, err, ErrOrderNotFound)
}

