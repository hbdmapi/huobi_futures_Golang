package index

type SubIndexKLineResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		Vol string `json:"vol"`

		Count string `json:"count"`

		Open string `json:"open"`

		Close string `json:"close"`

		Low string `json:"low"`

		High string `json:"high"`

		Amount string `json:"amount"`

		TradeTurnover string `json:"trade_turnover"`
	} `json:"tick"`
}
