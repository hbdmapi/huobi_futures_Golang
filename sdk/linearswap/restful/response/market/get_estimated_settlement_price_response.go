package market

type GetEstimatedSettlementPriceResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		EstimatedSettlementPrice string `json:"estimated_settlement_price"`

		SettlementType string `json:"settlement_type"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
