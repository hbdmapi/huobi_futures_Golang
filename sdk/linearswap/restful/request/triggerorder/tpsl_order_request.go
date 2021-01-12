package triggerorder

type TpslOrderRequest struct {
	ContractCode string `json:"contract_code"`

	Direction string `json:"direction"`

	Volume int64 `json:"volume"`

	TpTriggerPrice float32 `json:"tp_trigger_price"`

	TpOrderPrice float32 `json:"tp_order_price"`

	TpOrderPriceType string `json:"tp_order_price_type"`

	SlTriggerPrice float32 `json:"sl_trigger_price"`

	SlOrderPrice float32 `json:"sl_order_price"`

	SlOrderPriceType string `json:"sl_order_price_type"`
}

type BatchTpslOrderRequest []TpslOrderRequest
