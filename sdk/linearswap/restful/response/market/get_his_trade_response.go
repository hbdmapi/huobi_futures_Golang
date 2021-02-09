package market

type GetHisTradeResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	HisTrade []struct {
		Id int64 `json:"id"`

		Data []struct {
			Amount float32 `json:"amount"`

			Direction string `json:"direction"`

			Id int64 `json:"id"`

			Price float32 `json:"price"`

			Ts int64 `json:"ts"`

			Quantity float32 `json:"quantity"`

			TradeTurnover float32 `json:"trade_turnover"`
		} `json:"data"`

		Ts int64 `json:"ts"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
