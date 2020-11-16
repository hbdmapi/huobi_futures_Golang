package account

type GetApiTradingStatusResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	ts int64

	data struct {
		isDisable int `json:"is_disable"`

		orderPriceTypes string `json:"order_price_types"`

		disableReason string `json:"disable_reason"`

		disableInterval int64 `json:"disable_interval"`

		recoveryTime int64 `json:"recovery_time"`

		COR struct {
			ordersThreshold int64 `json:"orders_threshold"`

			orders int64

			invalidCancelOrders int64 `json:"invalid_cancel_orders"`

			cancelRatioThreshold float32 `json:"cancel_ratio_threshold"`

			cancelRatio float32 `json:"cancel_ratio"`

			isTrigger int `json:"cancel_ratio"`

			isActive int `json:"is_active"`
		}

		TDN struct {
			disablesThreshold int64 `json:"disables_threshold"`

			disables int64

			isTrigger int `json:"is_trigger"`

			isActive int `json:"is_active"`
		}
	} `json:"omitempty"`
}
