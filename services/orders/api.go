package orders

type CreateOrder struct {
	ID          string `json:"id"`
	Number      string `json:"number"`
	DeliveryId  string `json:"delivery_id"`
	RecipientId string `json:"recipient_id"`
}

type GetOrders struct {
	Orders []*CreateOrder `json:"orders"`
}

type DeleteOrder struct {
	ID string `json:"id"`
}
