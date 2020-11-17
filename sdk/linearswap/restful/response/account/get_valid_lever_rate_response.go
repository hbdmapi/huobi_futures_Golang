package account

type GetValidLeverRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		AvailableLeverRate string `json:"available_level_rate"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
