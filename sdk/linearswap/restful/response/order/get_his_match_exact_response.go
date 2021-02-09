package order

type GetHisMatchExactResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Trades []struct {
			Id string `json:"id"`

			MatchId int64 `json:"match_id"`

			QueryId int64 `json:"query_id"`

			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			TradeVolume float32 `json:"trade_volume"`

			TradePrice float32 `json:"trade_price"`

			TradeTurnover float32 `json:"trade_turnover"`

			CreateDate int64 `json:"create_date"`

			OffsetProfitloss float32 `json:"offset_profitloss"`

			RealProfit float32 `json:"real_profit"`

			TradeFee float32 `json:"trade_fee"`

			Role string `json:"role"`

			FeeAsset string `json:"fee_asset"`

			OrderSource string `json:"order_source"`
		} `json:"trades"`

		RemainSize int `json:"remain_size"`

		NextId int64 `json:"next_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
