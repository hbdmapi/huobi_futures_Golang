package market

type SubTradeDetailResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		Ts int64 `json:"ts"`

		Data []struct {
			Amount float32 `json:"amount"`

			Ts int64 `json:"ts"`

			Id int64 `json:"id"`

			Price float32 `json:"price"`

			Direction string `json:"direction"`

			Quantity float32 `json:"quantity"`

			TradeTurnover float32 `json:"trade_turnover"`
		} `json:"data"`
	} `json:"tick"`
}
