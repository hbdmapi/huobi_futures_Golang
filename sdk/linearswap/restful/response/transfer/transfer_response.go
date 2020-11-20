package transfer

type TransferResponse struct {
	Code int64 `json:"code"`

	Data int64 `json:"data",omitempty`

	Message string `json:"message"`

	Success bool `json:"success"`

	PrintLog bool `json:"print-log"`
}
