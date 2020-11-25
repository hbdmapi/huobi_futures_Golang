package market

type ReqKLineResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Data []struct {
		Id int64 `json:"id"`

		Vol float32 `json:"vol"`

		Count float32 `json:"count"`

		Open float32 `json:"open"`

		Close float32 `json:"close"`

		Low float32 `json:"low"`

		High float32 `json:"high"`

		Amount float32 `json:"amount"`

		TradeTurnover float32 `json:"trade_turnover"`
	} `json:"data"`
}
