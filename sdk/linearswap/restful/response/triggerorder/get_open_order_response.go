package triggerorder

type GetOpenOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Order []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			TriggerType string `json:"trigger_type"`

			Volume float32 `json:"volume"`

			OrderType int `json:"order_type"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			LeverRate int `json:"lever_rate"`

			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			OrderSource string `json:"order_source"`

			TriggerPrice float32 `json:"trigger_price"`

			OrderPrice float32 `json:"order_price"`

			CreatedAt int64 `json:"created_at"`

			OrderPriceType string `json:"order_price_type"`

			Status int `json:"status"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`

			TpslOrderType string `json:"tpsl_order_type,omitempty"`

			SourceOrderIßd string `json:"source_order_id,omitempty"`

			RelationTpslOrderId string `json:"relation_tpsl_order_id,omitempty"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
