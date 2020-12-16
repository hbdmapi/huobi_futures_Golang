package account

type GetSubAccountListResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		SubUid int64 `json:"sub_uid"`

		list []struct {
			Symbol string `json:"symbol"`

			ContractCode string `json:"contract_code"`

			MarginAsset string `json:"margin_asset"`

			MarginBalance float32 `json:"margin_balance"`

			LiquidationPrice float32 `json:"liquidation_price,omitempty"`

			RiskRate float32 `json:"risk_rate,omitempty"`

			MarginMode string `json:"margin_mode"`

			MarginAccount string `json:"margin_account"`
		}
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
