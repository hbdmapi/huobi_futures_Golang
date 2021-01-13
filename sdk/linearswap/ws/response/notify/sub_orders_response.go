package notify

type SubOrdersResponse struct {
	Op string `json:"op"`

	Topic string `json:"topic"`

	Ts int64 `json:"ts"`

	Uid string `json:"uid"`

	Symbol string `json:"symbol"`

	ContractCode string `json:"contract_code"`

	Volume float32 `json:"volume"`

	Price float32 `json:"price"`

	OrderPriceType string `json:"order_price_type"`

	Direction string `json:"direction"`

	Offset string `json:"offset"`

	Status int `json:"status"`

	LeverRate int `json:"lever_rate"`

	OrderId int64 `json:"order_id"`

	OrderIdStr string `json:"order_id_str"`

	ClientOrderId int64 `json:"client_order_id,omitempty"`

	OrderSource string `json:"order_source"`

	OrderType int `json:"order_type"`

	CreatedAt int64 `json:"created_at"`

	TradeVolume float32 `json:"trade_volume"`

	TradeTurnover float32 `json:"trade_turnover"`

	Fee float32 `json:"fee"`

	TradeAvgPrice float32 `json:"trade_avg_price"`

	MarginAsset string `json:"margin_asset"`

	MarginFrozen float32 `json:"margin_frozen"`

	Profit float32 `json:"profit"`

	LiquidationType float32 `json:"liquidation_type"`

	CanceledAt int64 `json:"canceled_at"`

	FeeAsset string `json:"fee_asset"`

	IsTpsl int `json:"is_tpsl"`

	Trade []struct {
		Id string `json:"id"`

		TradeId int64 `json:"trade_id"`

		TradeVolume float32 `json:"trade_volume"`

		TradePrice float32 `json:"trade_price"`

		TradeFee float32 `json:"trade_fee"`

		TradeTurnover float32 `json:"trade_turnover"`

		CreatedAt int64 `json:"created_at"`

		Role string `json:"role"`

		FeeAsset string `json:"fee_asset"`
	} `json:"trade"`
}
