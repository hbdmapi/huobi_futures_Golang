package notify

type SubContractInfoResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	EventSender string `json:"event"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ContractSize float32 `json:"contract_size"`

		PriceTick float32 `json:"price_tick"`

		SettlementDate string `json:"settlement_date"`

		CreateDate string `json:"create_date"`

		ContractStatus int `json:"contract_status"`
	} `json:"data"`
}
