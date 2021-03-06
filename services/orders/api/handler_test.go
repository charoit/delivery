package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"delivery/auth"
	"delivery/models"
	"delivery/services/orders"
	"delivery/services/orders/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	testUser := &models.User{
		ID: uuid.New().String(),
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	svc := new(service.ServiceMock)
	RegisterHTTPEndpoints(group, svc)

	inp := &orders.Order{
		Number:      "21-01-00001",
		DeliveryId:  uuid.New().String(),
		RecipientId: uuid.New().String(),
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	svc.On("Create", testUser, inp).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

func TestGet(t *testing.T) {
	testUser := &models.User{
		ID: uuid.New().String(),
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	svc := new(service.ServiceMock)
	RegisterHTTPEndpoints(group, svc)

	list := make([]*models.Order, 5)
	for i := 0; i < 5; i++ {
		list[i] = &models.Order{
			ID:        fmt.Sprintf("ID-%d", i),
			Number:    fmt.Sprintf("21-01-0000%d", i),
			Manager:   testUser,
			Date:      time.Now(),
			Delivery:  models.Address{},
			Recipient: models.Recipient{},
		}
	}

	svc.On("List", testUser).Return(list, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/orders", nil)
	r.ServeHTTP(w, req)

	expectedOut := svc.ToOrderList(list)
	expectedOutBody, err := json.Marshal(expectedOut)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(expectedOutBody), w.Body.String())
}

func TestRemove(t *testing.T) {
	testUser := &models.User{
		ID: uuid.New().String(),
	}

	r := gin.Default()
	group := r.Group("/api", func(c *gin.Context) {
		c.Set(auth.CtxUserKey, testUser)
	})

	svc := new(service.ServiceMock)
	RegisterHTTPEndpoints(group, svc)

	inp := &orders.Remove{
		ID: "ID-0",
	}

	body, err := json.Marshal(inp)
	assert.NoError(t, err)

	svc.On("Remove", testUser, inp).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/orders", bytes.NewBuffer(body))
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
