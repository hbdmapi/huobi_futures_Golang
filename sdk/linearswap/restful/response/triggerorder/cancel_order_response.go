package triggerorder

type CancelOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Error []struct {
			OrderId string `json:"order_id"`

			ErrorCode int `json:"err_code"`

			ErrorMessage string `json:"err_msg"`
		} `json:"errors"`

		Successes string `json:"successes"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
