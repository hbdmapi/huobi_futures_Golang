package market

type GetKLineResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

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
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
