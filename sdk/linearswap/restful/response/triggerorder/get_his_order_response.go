package triggerorder

type GetHisOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Orders []struct {
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

			RelationOrderId string `json:"relation_order_id"`

			OrderPriceType string `json:"order_price_type"`

			Status int `json:"status"`

			OrderSource string `json:"order_source"`

			TriggerPrice float32 `json:"trigger_price"`

			TriggeredPrice float32 `json:"triggered_price,omitempty"`

			OrderPrice float32 `json:"order_price"`

			CreatedAt int64 `json:"created_at"`

			TriggeredAt int64 `json:"triggered_at"`

			OrderInsertAt int64 `json:"order_insert_at"`

			CanceledAt int64 `json:"canceled_at,omitempty"`

			FailCode int `json:"fail_code,omitempty"`

			FailReason string `json:"fail_reason,omitempty"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
