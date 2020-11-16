package account

type GetOrderLimitResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data struct {
		orderPriceType string `json:"order_price_type"`

		list []struct {
			symbol string

			contractCode string `json:"contract_code"`

			openLimit float32 `json:"open_limit"`

			closeLimit float32 `json:"close_limit"`
		}
	} `json:"omitempty"`

	ts int64
}
