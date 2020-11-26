package notify

type SubLiquidationOrdersResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Direction string `json:"direction"`

		Offset string `json:"offset"`

		Volume float32 `json:"volume"`

		Price float32 `json:"price"`

		CreatedAt int64 `json:"created_at"`
	} `json:"data"`
}
