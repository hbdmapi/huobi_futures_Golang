package market

type GetEliteRatioResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Symbol string `json:"symbol"`

		ContractCode string `json:"contract_code"`

		ShortLongRatio []struct {
			BuyRatio float32 `json:"buy_ratio"`

			SellRatio float32 `json:"sell_ratio"`

			LockedRatio float32 `json:"locked_ratio,omitempty"`

			Ts int64 `json:"ts"`
		} `json:"list"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
