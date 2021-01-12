package triggerorder

type GetRelationTpslOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Volume float32 `json:"volume"`

		Price float32 `json:"price"`

		OrderPriceType string `json:"order_price_type"`

		Direction string `json:"direction"`

		Offset string `json:"offset"`

		LeverRate int `json:"lever_rate"`

		OrderId int64 `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`

		ClientOrderId int64 `json:"client_order_id,omitempty"`

		CreatedAt int64 `json:"created_at"`

		TradeVolume float32 `json:"trade_volume"`

		TradeTurnover float32 `json:"trade_turnover"`

		Fee float32 `json:"fee"`

		TradeAvgPrice float32 `json:"trade_avg_price"`

		//MarginAsset string `json:"margin_asset"`

		MarginFrozen float32 `json:"margin_frozen"`

		Profit float32 `json:"profit"`

		Status int `json:"status"`

		OrderType int `json:"order_type"`

		OrderSource string `json:"order_source"`

		FeeAsset string `json:"fee_asset"`

		//LiquidationType string `json:"liquidation_type"`

		CanceledAt int64 `json:"canceled_at"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		TpslOrderInfo []struct {
			Volume float32 `json:"volume"`

			TpslOrderType string `json:"tpsl_order_type,omitempty"`

			Direction string `json:"direction"`

			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			TriggerType string `json:"trigger_type"`

			TriggerPrice float32 `json:"trigger_price"`

			CreatedAt int64 `json:"created_at"`

			OrderPriceType string `json:"order_price_type"`

			Status int `json:"status"`

			RelationTpslOrderId string `json:"relation_tpsl_order_id,omitempty"`

			CanceledAt int64 `json:"canceled_at,omitempty"`

			FailCode int `json:"fail_code,omitempty"`

			FailReason string `json:"fail_reason,omitempty"`

			TriggeredPrice float32 `json:"triggered_price,omitempty"`

			RelationOrderId string `json:"relation_order_id"`
		} `json:"tpsl_order_info"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
