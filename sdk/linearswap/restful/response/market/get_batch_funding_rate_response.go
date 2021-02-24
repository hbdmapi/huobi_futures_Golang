package market

type GetBatchFundingRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		FeeAsset string `json:"fee_asset"`

		FundingTime string `json:"funding_time"`

		FundingRate string `json:"funding_rate"`

		EstimatedRate string `json:"estimated_rate"`

		NextFundingTime string `json:"next_funding_time"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
