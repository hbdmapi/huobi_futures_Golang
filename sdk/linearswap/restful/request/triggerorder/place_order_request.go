package triggerorder

type PlaceOrderRequest struct {
	ContractCode string `json:"contract_code"`

	TriggerType string `json:"trigger_type"`

	TriggerPrice float32 `json:"trigger_price"`

	OrderPrice float32 `json:"order_price,omitempty"`

	OrderPriceType string `json:"order_price_type,omitempty"`

	Volume int64 `json:"volume"`

	Direction string `json:"direction"`

	Offset string `json:"offset"`

	LeverRate int `json:"lever_rate,omitempty"`
}
