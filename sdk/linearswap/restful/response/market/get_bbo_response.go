package market

type GetBboResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Ticks []struct {
		ContractCode string `json:"contract_code"`

		Mrid int64 `json:"mrid"`

		Ask []float32 `json:"ask"`

		Bid []float32 `json:"bid"`

		Ts int64 `json:"ts"`
	} `json:"ticks,omitempty"`

	Ts int64 `json:"ts"`
}
