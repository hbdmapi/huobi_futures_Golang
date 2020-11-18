package market

type GetMergedResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Id int64 `json:"id"`

		Amount string `json:"status"`

		Ask []float32 `json:"ask"`

		Bid []float32 `json:"bid"`

		Open string `json:"open"`

		Close string `json:"close"`

		Count float32 `json:"count"`

		High string `json:"high"`

		Low string `json:"low"`

		Vol string `json:"vol"`

		TradeTurnover string `json:"trade_turnover"`

		Ts int64 `json:"ts"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}
