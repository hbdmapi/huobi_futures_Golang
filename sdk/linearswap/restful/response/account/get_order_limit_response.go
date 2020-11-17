package account

type GetOrderLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderPriceType string `json:"order_price_type"`

		List []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			OpenLimit float32 `json:"open_limit"`

			CloseLimit float32 `json:"close_limit"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
