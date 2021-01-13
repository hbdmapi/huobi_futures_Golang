package order

type GetOrderInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
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

		MarginAsset string `json:"margin_asset"`

		MarginFrozen float32 `json:"margin_frozen"`

		Profit float32 `json:"profit"`

		Status int `json:"status"`

		OrderType int `json:"order_type"`

		OrderSource string `json:"order_source"`

		FeeAsset string `json:"fee_asset"`

		LiquidationType string `json:"liquidation_type"`

		CanceledAt int64 `json:"canceled_at"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		IsTpsl int `json:"is_tpsl"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
