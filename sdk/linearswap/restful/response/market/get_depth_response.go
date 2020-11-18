package market

type GetDepthResponse struct {
	Ch string `json:"ch"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Tick struct {
		Asks [][]float32 `json:"asks"`

		Bids [][]float32 `json:"bids"`

		Ch string `json:"ch"`

		Id int64 `json:"id"`

		Mrid int64 `json:"mrid"`

		Ts int64 `json:"ts"`

		Version int64 `json:"version"`
	} `json:"tick,omitempty"`

	Ts int64 `json:"ts"`
}
