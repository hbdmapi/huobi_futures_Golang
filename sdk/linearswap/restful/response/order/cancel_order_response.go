package order

type CancelOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Error []struct {
			OrderId string `json:"order_id"`

			ErrorCode int `json:"err_code"`

			ErrorMessage string `json:"err_msg"`
		} `json:"errors,omitempty"`

		Successes string `json:"successes,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
