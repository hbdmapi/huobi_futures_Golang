package market

type GetStrKLineResponse struct {
	Ch string `json:"ch,omitempty"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

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
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
