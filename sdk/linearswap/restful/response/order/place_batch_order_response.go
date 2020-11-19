package order

type PlaceBatchOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Errors []struct {
			Index int `json:"index"`

			ErrorCode int `json:"err_code"`

			ErrorMessage string `json:"err_msg"`
		} `json:"errors,omitempty"`

		Success []struct {
			Index int `json:"index"`

			OrderId int64 `json:"order_id"`

			ClientOrderId int64 `json:"client_order_id,omitempty"`

			OrderIdStr string `json:"order_id_str"`
		} `json:"success,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
