package triggerorder

type PlaceOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId int `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
