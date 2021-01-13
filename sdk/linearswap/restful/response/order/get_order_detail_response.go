package order

type GetOrderDetailResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		LeverRate int `json:"lever_rate"`

		Direction string `json:"direction"`

		Offset string `json:"offset"`

		Volume float32 `json:"volume"`

		Price float32 `json:"price"`

		CreatedAt int64 `json:"created_at"`

		CanceledAt int64 `json:"canceled_at"`

		OrderSource string `json:"order_source"`

		OrderPriceType string `json:"order_price_type"`

		MarginAsset string `json:"margin_asset"`

		MarginFrozen float32 `json:"margin_frozen"`

		Profit float32 `json:"profit"`

		InstrumentPrice float32 `json:"instrument_price"`

		FinalInterest float32 `json:"final_interest"`

		AdjustValue float32 `json:"adjust_value"`

		Fee float32 `json:"fee"`

		FeeAsset string `json:"fee_asset"`

		LiquidationType string `json:"liquidation_type"`

		OrderId int64 `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`

		ClientOrderId int64 `json:"client_order_id,omitempty"`

		OrderType string `json:"order_type"`

		Status int `json:"status"`

		TradeAvgPrice float32 `json:"trade_avg_price,omitempty"`

		TradeTurnover float32 `json:"trade_turnover"`

		TradeVolume float32 `json:"trade_volume"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		IsTpsl int `json:"is_tpsl"`

		Trades []struct {
			Id string `json:"id"`

			TradeId int64 `json:"trade_id"`

			TradePrice float32 `json:"trade_price"`

			TradeVolume float32 `json:"trade_volume"`

			TradeTurnover float32 `json:"trade_turnover"`

			TradeFee float32 `json:"trade_fee"`

			Role string `json:"role"`

			CreatedAt int64 `json:"created_at"`
		} `json:"trades,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
