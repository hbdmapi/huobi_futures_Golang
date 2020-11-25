package index

type ReqIndexKLineResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Ts int64 `json:"ts"`

	Data []struct {
		Id int64 `json:"id"`

		Vol string `json:"vol"`

		Count string `json:"count"`

		Open string `json:"open"`

		Close string `json:"close"`

		Low string `json:"low"`

		High string `json:"high"`

		Amount string `json:"amount"`

		TradeTurnover string `json:"trade_turnover"`
	} `json:"data"`
}
