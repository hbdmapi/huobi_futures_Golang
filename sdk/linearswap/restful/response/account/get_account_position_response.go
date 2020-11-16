package account

type GetAccountPositionResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		symbol string

		contractCode string `json:"contract_code"`

		volume float32

		available float32

		frozen float32

		costOpen float32 `json:"cost_open"`

		costHold float32 `json:"cost_hold"`

		profitUnreal float32 `json:"profit_unreal"`

		profitRate float32 `json:"profit_rate"`

		profit float32

		marginAsset string `json:"margin_asset"`

		positionMargin float32 `json:"position_margin"`

		leverRate int `json:"lever_rate"`

		direction string

		lastPrice float32 `json:"last_price"`
	} `json:"omitempty"`

	ts int64
}
