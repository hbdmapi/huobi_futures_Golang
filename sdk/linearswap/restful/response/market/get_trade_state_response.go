package market

type GetTradeStateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		Open int `json:"open"`

		Close int `json:"close"`

		Cancel int `json:"cancel"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
