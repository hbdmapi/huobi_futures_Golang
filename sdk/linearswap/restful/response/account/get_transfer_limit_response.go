package account

type GetTransferLimitResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		symbol string

		contractCode string `json:"contract_code"`

		openLimit string `json:"open_limit"`

		transferInMaxEach float32 `json:"transfer_in_max_each"`

		transferInMinEach float32 `json:"transfer_in_min_each"`

		transferOutMaxEach float32 `json:"transfer_out_max_each"`

		transferOutMinEach float32 `json:"transfer_out_min_each"`

		transferInMaxDaily float32 `json:"transfer_in_max_daily"`

		transferOutMaxDaily float32 `json:"transfer_out_max_daily"`

		netTransferInMaxDaily float32 `json:"net_transfer_in_max_daily"`

		netTransferOutMaxDaily float32 `json:"net_transfer_out_max_daily"`
	} `json:"omitempty"`

	ts int64
}
