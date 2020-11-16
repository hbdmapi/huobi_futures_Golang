package account

type GetPositionLimitResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		symbol string

		contractCode string `json:"contract_code"`

		buyLimit float32 `json:"buy_limit"`

		sellLimit float32 `json:"sell_limit"`
	} `json:"omitempty"`

	ts int64
}
