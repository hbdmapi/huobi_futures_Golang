package triggerorder

type TpslOrderResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		TpOrder struct {
			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`
		} `json:"tp_order"`

		SlOrder struct {
			OrderId int64 `json:"order_id"`

			OrderIdStr string `json:"order_id_str"`
		} `json:"sl_order"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
