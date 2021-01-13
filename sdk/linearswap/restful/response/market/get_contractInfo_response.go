package market

type GetContractInfoResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ContractSize float32 `json:"contract_size"`

		PriceTick float32 `json:"price_tick"`

		SettlementDate string `json:"settlement_date"`

		CreateDate string `json:"create_date"`

		DeliveryTime string `json:"delivery_time"`

		ContractStatus int `json:"contract_status"`

		SupportMarginMode string `json:"support_margin_mode"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
