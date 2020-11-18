package market

type GetApiStatusResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Open int `json:"open"`

		Close int `json:"close"`

		Cancel int `json:"cancel"`

		TransferIn int `json:"transfer_in"`

		TransferOut int `json:"transfer_out"`

		MasterTransferSub int `json:"master_transfer_sub"`

		SubTransferMaster int `json:"sub_transfer_master"`

		MasterTransferSubInnerIn int `json:"master_transfer_sub_inner_in"`

		MasterTransferSubInnerOut int `json:"master_transfer_sub_inner_out"`

		SubTransferMasterInnerIn int `json:"sub_transfer_master_inner_in"`

		SubTransferMasterInnerOut int `json:"sub_transfer_master_inner_out"`

		TransferInnerIn int `json:"transfer_inner_in"`

		TransferInnerOut int `json:"transfer_inner_out"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
