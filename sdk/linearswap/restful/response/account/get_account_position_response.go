package account

type GetAccountPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		Volume float32 `json:"volume"`

		Available float32 `json:"available"`

		Frozen float32 `json:"frozen"`

		CostOpen float32 `json:"cost_open"`

		CostHold float32 `json:"cost_hold"`

		ProfitUnreal float32 `json:"profit_unreal"`

		ProfitRate float32 `json:"profit_rate"`

		Profit float32 `json:"profit"`

		MarginAsset string `json:"margin_asset"`

		PositionMargin float32 `json:"position_margin"`

		LeverRate int `json:"lever_rate"`

		Direction string `json:"direction"`

		LastPrice float32 `json:"last_price"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
