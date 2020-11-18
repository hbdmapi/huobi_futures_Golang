package market

type GetAdjustFactorFundResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		List []struct {
			LeverRate float32 `json:"lever_rate"`

			Ladder []struct {
				MinSize float32 `json:"min_size"`

				MaxSize float32 `json:"max_size,omitempty"`

				Ladder int `json:"ladder"`

				AdjustFactor float32 `json:"adjust_factor"`
			} `json:"ladders"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
