package notify

type SubPositionsResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	EventSender string `json:"event"`

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

		PositionMargin float32 `json:"position_margin"`

		LeverRate int `json:"lever_rate"`

		Direction string `json:"direction"`

		LastPrice float32 `json:"last_price"`

		MarginAsset string `json:"margin_asset"`

		MarginAccount string `json:"margin_account"`

		MarginMode string `json:"margin_mode"`
	} `json:"data"`
}
