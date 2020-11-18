package market

type GetHisOpenInterestResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Tick []struct {
			Volume float32 `json:"volume"`

			AmountType int `json:"amount_type"`

			Value float32 `json:"value"`

			Ts int64 `json:"ts"`
		} `json:"tick"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
