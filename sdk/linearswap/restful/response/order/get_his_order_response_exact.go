package order

type GetHisOrderExactResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Orders []struct {
			QueryId int64 `json:"query_id"`

			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`

			LeverRate int `json:"lever_rate"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			Volume float32 `json:"volume"`

			Price float32 `json:"price"`

			CreateDate int64 `json:"create_date"`

			OrderSource string `json:"order_source"`

			OrderPriceType string `json:"order_price_type"`

			MarginFrozen float32 `json:"margin_frozen"`

			Profit float32 `json:"profit"`

			RealProfit float32 `json:"real_profit"`

			TradeVolume float32 `json:"trade_volume"`

			TradeTurnover float32 `json:"trade_turnover"`

			Fee float32 `json:"fee"`

			TradeAvgPrice float32 `json:"trade_avg_price"`

			Status int `json:"status"`

			OrderType int `json:"order_type"`

			FeeAsset string `json:"fee_asset"`

			LiquidationType string `json:"liquidation_type"`

			IsTpsl int `json:"is_tpsl"`
		} `json:"orders"`

		RemainSize int `json:"remain_size"`

		NextId int64 `json:"next_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
