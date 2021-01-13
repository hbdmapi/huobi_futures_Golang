package order

type GetHisOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Orders []struct {
			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			LeverRate int `json:"lever_rate"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			Volume float32 `json:"volume"`

			Price float32 `json:"price"`

			CreateDate int64 `json:"create_date"`

			UpdateTime int64 `json:"update_time	"`

			OrderSource string `json:"order_source"`

			OrderPriceType int `json:"order_price_type"`

			MarginAsset string `json:"margin_asset"`

			MarginFrozen float32 `json:"margin_frozen"`

			Profit float32 `json:"profit"`

			TradeVolume float32 `json:"trade_volume"`

			TradeTurnover float32 `json:"trade_turnover"`

			Fee float32 `json:"fee"`

			TradeAvgPrice float32 `json:"trade_avg_price"`

			Status int `json:"status"`

			OrderType int `json:"order_type"`

			FeeAsset string `json:"fee_asset"`

			LiquidationType string `json:"liquidation_type"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`

			IsTpsl int `json:"is_tpsl"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
