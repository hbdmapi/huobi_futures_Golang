package account

type GetValidLeverRateResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		contractCode string `json:"contract_code"`

		availableLeverRate string `json:"available_level_rate"`
	} `json:"omitempty"`

	ts int64
}
