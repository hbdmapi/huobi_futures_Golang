package order

type GetHisMatchResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Trades []struct {
			Id string `json:"id"`

			MatchId int64 `json:"match_id"`

			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`

			Symbol string `json:"symbol"`

			OrderSource string `json:"order_source"`

			ContractCode string `json:"contract_code"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			TradeVolume float32 `json:"trade_volume"`

			TradePrice float32 `json:"trade_price"`

			TradeTurnover float32 `json:"trade_turnover"`

			CreateDate int64 `json:"create_date"`

			OffsetProfitloss float32 `json:"offset_profitloss"`

			TradeFee float32 `json:"trade_fee"`

			Role string `json:"role"`

			FeeAsset string `json:"fee_asset"`
		} `json:"trades"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
