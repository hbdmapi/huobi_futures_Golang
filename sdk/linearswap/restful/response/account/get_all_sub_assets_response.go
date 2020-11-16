package account

type GetAllSubAssetsResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		subUid string `json:"sub_uid"`

		list []struct {
			symbol string

			contractCode string `json:"contract_code"`

			marginAsset string `json:"margin_asset"`

			marginBalance float32 `json:"margin_balance"`

			liquidationPrice float32 `json:"liquidation_price,omitempty"`

			riskRate float32 `json:"risk_rate,omitempty"`
		}
	} `json:"omitempty"`

	ts int64
}
