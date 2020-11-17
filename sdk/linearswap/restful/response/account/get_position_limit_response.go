package account

type GetPositionLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		BuyLimit float32 `json:"buy_limit"`

		SellLimit float32 `json:"sell_limit"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
