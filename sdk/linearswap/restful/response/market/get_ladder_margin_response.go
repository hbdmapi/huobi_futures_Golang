package market

type GetLadderMarginResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		List []struct {
			LeverRate int `json:"lever_rate"`

			Ladders []struct {
				MinMarginBalance float32 `json:"min_margin_balance,omitempty"`

				MaxMarginBalance float32 `json:"max_margin_balance,omitempty"`

				MinMarginAvailable float32 `json:"min_margin_available,omitempty"`

				MaxMarginAvailable float32 `json:"max_margin_available,omitempty"`
			} `json:"ladders"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
