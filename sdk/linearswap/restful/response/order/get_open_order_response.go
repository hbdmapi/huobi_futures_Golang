package order

type GetOpenOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Orders []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			Volume float32 `json:"volume"`

			Price float32 `json:"price"`

			OrderPriceType string `json:"order_price_type"`

			OrderType int `json:"order_type"`

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

			FeeAsset string `json:"fee_asset"`

			TradeAvgPrice float32 `json:"trade_avg_price,omitempty"`

			MarginFrozen float32 `json:"margin_frozen"`

			MarginAsset string `json:"margin_asset"`

			Profit float32 `json:"profit"`

			Status int `json:"status"`

			OrderSource string `json:"order_source"`

			LiquidationType string `json:"liquidation_type"`

			CanceledAt int64 `json:"canceled_at,omitempty"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
