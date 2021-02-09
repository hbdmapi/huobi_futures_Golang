package market

type ReqTradeDetailResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Ts int64 `json:"ts"`

	Data []struct {
		Id int64 `json:"id"`

		Price string `json:"price"`

		Amount string `json:"amount"`

		Direction string `json:"direction"`

		Ts int64 `json:"ts"`

		Quantity string `json:"quantity"`

		TradeTurnover string `json:"trade_turnover"`
	} `json:"data"`
}
