package market

type SubBBOResponse struct {
	Ch string `json:"ch"`

	Ts int64 `json:"ts"`

	Tick struct {
		Ch string `json:"ch"`

		Mrid int64 `json:"mrid"`

		Id int64 `json:"id"`

		Ask []float32 `json:"ask"`

		Did []float32 `json:"bid"`

		Version int64 `json:"version"`

		Ts int64 `json:"ts"`
	} `json:"tick"`
}
