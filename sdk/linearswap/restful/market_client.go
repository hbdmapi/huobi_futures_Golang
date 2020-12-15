package restful

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/restful/response/market"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
)

type MarketClient struct {
	PUrlBuilder *reqbuilder.PublicUrlBuilder
}

func (mc *MarketClient) Init(host string) *MarketClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	mc.PUrlBuilder = new(reqbuilder.PublicUrlBuilder).Init(host)
	return mc
}

func (mc *MarketClient) GetContractInfoAsync(data chan market.GetContractInfoResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_contract_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetContractInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetContractInfoAsync error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetIndexAsync(data chan market.GetIndexResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_index"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetIndexResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetIndexResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetPriceLimitAsync(data chan market.GetPriceLimitResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_price_limit?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetPriceLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetOpenInterestAsync(data chan market.GetOpenInterestResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_open_interest"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetDepthAsync(data chan market.GetDepthResponse, contractCode string, fcType string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/depth?contract_code=%s&&type=%s", contractCode, fcType)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetDepthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetDepthResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetKLineAsync(data chan market.GetKLineResponse, contractCode string, period string, size int, from int, to int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/kline?contract_code=%s&&period=%s", contractCode, period)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&&size=%d", size)
	}
	if from != 0 {
		option += fmt.Sprintf("&&from=%d", from)
	}
	if to != 0 {
		option += fmt.Sprintf("&&to=%d", to)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetKLineResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetMergedAsync(data chan market.GetMergedResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/detail/merged?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetMergedResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetMergedResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetTradeAsync(data chan market.GetTradeResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/trade?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTradeResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisTradeAsync(data chan market.GetHisTradeResponse, contractCode string, size int) {
	// location
	location := fmt.Sprintf("/linear-swap-ex/market/history/trade?contract_code=%s&&size=%d", contractCode, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisTradeResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetRiskInfoAsync(data chan market.GetRiskInfoResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_risk_info"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("contract_code=%s", contractCode)
	}
	if option != "" {
		location += fmt.Sprintf("?%s", option)
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetRiskInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetRiskInfoResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetInsuranceFundAsync(data chan market.GetInsuranceFundResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_insurance_fund?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&&size=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&&from=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetInsuranceFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetInsuranceFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) IsolatedGetAdjustFactorFundAsync(data chan market.GetAdjustFactorFundResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetAdjustFactorFundAsync(data chan market.GetAdjustFactorFundResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_adjustfactor"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetAdjustFactorFundResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAdjustFactorFundResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisOpenInterestAsync(data chan market.GetHisOpenInterestResponse, contractCode string, period string, amountType int, size int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_his_open_interest?contract_code=%s&&period=%s&&amount_type=%d", contractCode, period, amountType)

	// option
	option := ""
	if size != 0 {
		option += fmt.Sprintf("&&size=%d", size)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisOpenInterestResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOpenInterestResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetEliteAccountRatioAsync(data chan market.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_account_ratio?contract_code=%s&&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetElitePositionRatioAsync(data chan market.GetEliteRatioResponse, contractCode string, period string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_elite_position_ratio?contract_code=%s&&period=%s", contractCode, period)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetEliteRatioResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetElitePositionRatioResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) IsolatedGetApiStateAsync(data chan market.GetApiStateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_api_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetApiStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetApiStateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetTransferStateAsync(data chan market.GetTransferStateResponse, marginAccount string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_transfer_state"

	// option
	if marginAccount == "" {
		marginAccount = "USDT"
	}
	location += fmt.Sprintf("?margin_account=%s", marginAccount)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTransferStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferStateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) CrossGetTradeStateAsync(data chan market.GetTradeStateResponse, contractCode string) {
	// location
	location := "/linear-swap-api/v1/swap_cross_trade_state"

	// option
	option := ""
	if contractCode != "" {
		option += fmt.Sprintf("?contract_code=%s", contractCode)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetTradeStateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTradeStatusResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetFundingRateAsync(data chan market.GetFundingRateResponse, contractCode string) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_funding_rate?contract_code=%s", contractCode)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetHisFundingRateAsync(data chan market.GetHisFundingRateResponse, contractCode string, pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_historical_funding_rate?contract_code=%s", contractCode)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetHisFundingRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisFundingRateResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetLiquidationOrdersAsync(data chan market.GetLiquidationOrdersResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// location
	location := fmt.Sprintf("/linear-swap-api/v1/swap_liquidation_orders?contract_code=%s&&trade_type=%d&&create_date=%d", contractCode, tradeType, createDate)

	// option
	option := ""
	if pageIndex != 0 {
		option += fmt.Sprintf("&&page_index=%d", pageIndex)
	}
	if pageSize != 0 {
		option += fmt.Sprintf("&&page_size=%d", pageSize)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetLiquidationOrdersResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetLiquidationOrdersResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetPremiumIndexKLineAsync(data chan market.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_premium_index_kline?contract_code=%s&&period=%s&&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPriceLimitResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetEstimatedRateKLineAsync(data chan market.GetStrKLineResponse, contractCode string, period string, size int) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_estimated_rate_kline?contract_code=%s&&period=%s&&size=%d", contractCode, period, size)

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetStrKLineResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetStrKLineResponse error: %s", getErr)
	}
	data <- result
}

func (mc *MarketClient) GetBasisAsync(data chan market.GetBasisResponse, contractCode string, period string, size int, basisPriceType string) {
	// location
	location := fmt.Sprintf("/index/market/history/linear_swap_basis?contract_code=%s&&period=%s&&size=%d", contractCode, period, size)

	// option
	option := ""
	if basisPriceType != "" {
		option += fmt.Sprintf("&&basis_price_type=%s", basisPriceType)
	}
	if option != "" {
		location += option
	}

	url := mc.PUrlBuilder.Build(location, nil)
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := market.GetBasisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBasisResponse error: %s", getErr)
	}
	data <- result
}
