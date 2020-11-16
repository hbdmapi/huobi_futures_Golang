package account

type GetAccountAssetsResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		symbol string

		contractCode string `json:"contract_code"`

		marginAsset string `json:"margin_asset"`

		marginBalance float32 `json:"margin_balance"`

		marginStatic float32 `json:"margin_static"`

		marginPosition float32 `json:"margin_position"`

		marginFrozen float32 `json:"margin_frozen"`

		marginAvailable float32 `json:"margin_available"`

		profitReal float32 `json:"profit_real"`

		profitUnreal float32 `json:"profit_unreal"`

		withdrawAvailable float32 `json:"withdraw_available,omitempty"`

		riskRate float32 `json:"risk_rate,omitempty"`

		liquidationPrice float32 `json:"liquidation_price,omitempty"`

		leverRate float32 `json:"lever_rate"`

		adjustFactor float32 `json:"adjust_factor"`
	} `json:"omitempty"`

	ts int64
}
