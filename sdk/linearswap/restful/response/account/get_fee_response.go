package account

type GetFeeResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		OpenMakerFee string `json:"open_maker_fee"`

		OpenTakerFee string `json:"open_taker_fee"`

		CloseMakerFee string `json:"close_maker_fee"`

		CloseTakerFee string `json:"close_taker_fee"`

		FeeAsset string `json:"fee_asset"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
