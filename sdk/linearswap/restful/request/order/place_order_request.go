package order

type PlaceOrderRequest struct {
	ContractCode string `json:"contract_code"`

	ClientOrderId int64 `json:"client_order_id,omitempty"`

	Price float32 `json:"price"`

	Volume int64 `json:"volume"`

	Direction string `json:"direction"`

	Offset string `json:"offset"`

	LeverRate int `json:"lever_rate"`

	OrderPriceType string `json:"order_price_type"`
}

type BatchPlaceOrderRequest []PlaceOrderRequest
