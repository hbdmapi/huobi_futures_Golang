package market

type GetRiskInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		EstimatedClawback float32 `json:"estimated_clawback"`

		InsuranceFund float32 `json:"insurance_fund"`

		ContractCode string `json:"contract_code"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
