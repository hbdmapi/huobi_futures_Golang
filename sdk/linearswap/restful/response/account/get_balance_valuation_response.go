package account

type GetBalanceValuationResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		ValuationAsset string `json:"valuation_asset"`

		Balance string `json:"balance"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
