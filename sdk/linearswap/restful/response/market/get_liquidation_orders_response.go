package market

type GetLiquidationOrdersResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Order []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			CreatedAt int64 `json:"created_at"`

			Direction string `json:"direction"`

			Offset string `json:"offset"`

			Price float32 `json:"price"`

			Volume float32 `json:"volume"`
		} `json:"orders"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
