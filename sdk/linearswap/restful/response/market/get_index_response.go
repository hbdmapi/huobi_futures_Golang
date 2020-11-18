package market

type GetIndexResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ContractCode string `json:"contract_code"`

		IndexPrice float32 `json:"index_price"`

		IndexTs int64 `json:"index_ts"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
