package notify

type SubFundingRateResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		FeeAsset string `json:"fee_asset"`

		FundingTime string `json:"funding_time"`

		FundingRate string `json:"funding_rate"`

		EstimatedRate string `json:"estimated_rate"`

		SettlementTime string `json:"settlement_time"`
	} `json:"data"`
}
