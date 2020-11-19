package order

type LightningCloseResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		OrderId int64 `json:"order_id"`

		OrderIdStr string `json:"order_id_str"`

		ClientOrderId int64 `json:"client_order_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
