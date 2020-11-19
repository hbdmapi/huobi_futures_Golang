package order

type PlaceOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId int64 `json:"order_id"`

		ClientOrderId int64 `json:"client_order_id,omitempty"`

		OrderIdStr string `json:"order_id_str"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
