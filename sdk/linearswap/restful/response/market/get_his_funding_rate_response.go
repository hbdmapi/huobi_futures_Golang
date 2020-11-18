package market

type GetHisFundingRateResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Data []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			FeeAsset string `json:"fee_asset"`

			FundingTime string `json:"funding_time"`

			FundingRate string `json:"funding_rate"`

			RealizedRate string `json:"realized_rate"`

			AvgPremiumIndex string `json:"avg_premium_index"`
		} `json:"data"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
