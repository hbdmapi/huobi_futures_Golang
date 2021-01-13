package market

type SubKLineResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		Mrid int64 `json:"mrid"`

		Vol float32 `json:"vol"`

		Count float32 `json:"count"`

		Open float32 `json:"open"`

		Close float32 `json:"close"`

		Low float32 `json:"low"`

		High float32 `json:"high"`

		Amount float32 `json:"amount"`

		TradeTurnover float32 `json:"trade_turnover"`

		Ask []float32 `json:"ask,omitempty"`

		Bid []float32 `json:"bid,omitempty"`
	} `json:"tick"`
}
