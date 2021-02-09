package account

type GetSubAccountInfoListResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		SubList []struct {
			SubUid int64 `json:"sub_uid"`

			AccountInfoList []struct {
				Symbol string `json:"symbol"`

				ContractCode string `json:"contract_code"`

				MarginMode string `json:"margin_mode"`

				MarginAccount string `json:"margin_account"`

				MarginAsset string `json:"margin_asset"`

				MarginBalance float32 `json:"margin_balance"`

				LiquidationPrice float32 `json:"liquidation_price,omitempty"`

				RiskRate float32 `json:"risk_rate,omitempty"`
			} `json:"account_info_list"`
		} `json:"sub_list"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
