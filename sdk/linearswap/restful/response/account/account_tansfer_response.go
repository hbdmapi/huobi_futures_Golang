package account

type AccountTransferResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data struct {
		orderId string `json:"order_id"`
	} `json:"omitempty"`

	ts int64
}
