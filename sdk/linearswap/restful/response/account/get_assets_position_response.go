package account

type GetAssetsPositionResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Symbol string `json:"symbol,omitempty"`

		ContractCode string `json:"contract_code,omitempty"`

		MarginAsset string `json:"margin_asset"`

		MarginBalance float32 `json:"margin_balance"`

		MarginStatic float32 `json:"margin_static"`

		MarginPosition float32 `json:"margin_position"`

		MarginFrozen float32 `json:"margin_frozen"`

		MarginAvailable float32 `json:"margin_available,omitempty"`

		ProfitReal float32 `json:"profit_real"`

		ProfitUnreal float32 `json:"profit_unreal"`

		WithdrawAvailable float32 `json:"withdraw_available,omitempty"`

		RiskRate float32 `json:"risk_rate,omitempty"`

		LiquidationPrice float32 `json:"liquidation_price,omitempty"`

		LeverRate float32 `json:"lever_rate,omitempty"`

		AdjustFactor float32 `json:"adjust_factor,omitempty"`

		MarginMode string `json:"margin_mode"`

		MarginAccount string `json:"margin_account"`

		ContractDetail []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginPosition float32 `json:"margin_position"`

			MarginFrozen float32 `json:"margin_frozen"`

			MarginAvailable float32 `json:"margin_available"`

			ProfitUnreal float32 `json:"profit_unreal"`

			LiquidationPrice float32 `json:"liquidation_price"`

			LeverRate float32 `json:"lever_rate"`

			AdjustFactor float32 `json:"adjust_factor"`
		} `json:"contract_detail,omitempty"`

		Positions []struct {
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

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		} `json:"positions,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
