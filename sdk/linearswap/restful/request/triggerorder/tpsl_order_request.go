package triggerorder

type TpslOrderRequest struct {
	ContractCode string `json:"contract_code"`

	Direction string `json:"direction"`

	Volume int64 `json:"volume"`

	TpTriggerPrice float32 `json:"tp_trigger_price,omitempty"`

	TpOrderPrice float32 `json:"tp_order_price,omitempty"`

	TpOrderPriceType string `json:"tp_order_price_type,omitempty"`

	SlTriggerPrice float32 `json:"sl_trigger_price,omitempty"`

	SlOrderPrice float32 `json:"sl_order_price,omitempty"`

	SlOrderPriceType string `json:"sl_order_price_type,omitempty"`
}

type BatchTpslOrderRequest []TpslOrderRequest
