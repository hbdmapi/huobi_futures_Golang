package market

type SubDepthResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Mrid int64 `json:"mrid"`

		Id int64 `json:"id"`

		Asks [][]float32 `json:"asks"`

		Bids [][]float32 `json:"bids"`

		Ts int64 `json:"ts"`

		Version int64 `json:"version"`

		Ch string `json:"ch"`

		TickEvent string `json:"event,omitempty"`
	} `json:"tick"`
}
