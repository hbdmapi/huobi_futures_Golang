package account

type GetTransferLimitResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		OpenLimit string `json:"open_limit"`

		TransferInMaxEach float32 `json:"transfer_in_max_each"`

		TransferInMinEach float32 `json:"transfer_in_min_each"`

		TransferOutMaxEach float32 `json:"transfer_out_max_each"`

		TransferOutMinEach float32 `json:"transfer_out_min_each"`

		TransferInMaxDaily float32 `json:"transfer_in_max_daily"`

		TransferOutMaxDaily float32 `json:"transfer_out_max_daily"`

		NetTransferInMaxDaily float32 `json:"net_transfer_in_max_daily"`

		NetTransferOutMaxDaily float32 `json:"net_transfer_out_max_daily"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
