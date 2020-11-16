package account

type GetFeeResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		symbol string

		contractCode string `json:"contract_code"`

		openMakerFee string `json:"open_maker_fee"`

		openTakerFee string `json:"open_taker_fee"`

		closeMakerFee string `json:"close_maker_fee"`

		closeTakerFee string `json:"close_taker_fee"`

		feeAsset string `json:"fee_asset"`
	} `json:"omitempty"`

	ts int64
}
