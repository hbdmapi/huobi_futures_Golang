package restful

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
)

type AccountClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (ac *AccountClient) Init(accessKey string, secretKey string, host string) *AccountClient {
	if host == "" {
		host = linearswap.DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return ac
}

func (ac *AccountClient) GetAccountAssetsAsync(data chan account.GetAccountAssetsResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountAssetsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountAssetsResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetAccountPositionAsync(data chan account.GetAccountPositionResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info", nil)
	}

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d", subUid)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountPositionResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetAllSubAssetsAsync(data chan account.GetAllSubAssetsResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAllSubAssetsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAllSubAssetsResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) AccountTransferAsync(data chan account.AccountTransferResponse, asset string, fromMarginAccount string, toMarginAccount string, amount float64,
	subUid int64, fcType string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer", nil)
	if subUid == 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_inner", nil)
	}

	// content
	content := fmt.Sprintf(",\"asset\":\"%s\", \"from_margin_account\":\"%s\", \"to_margin_account\":\"%s\", \"amount\":%f",
		asset, fromMarginAccount, toMarginAccount, amount)
	if subUid != 0 {
		content += fmt.Sprintf(",\"sub_uid\": %d,\"type\": \"%s\"", subUid, fcType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.AccountTransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to AccountTransferResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetAccountTransHisAsync(data chan account.GetAccountTransHisResponse, marginAccount string,
	beMasterSub bool, fcType string, createDate int, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record", nil)
	if beMasterSub {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_master_sub_transfer_record", nil)
	}

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
		if beMasterSub {
			content += fmt.Sprintf(",\"transfer_type\": \"%s\"", fcType)
		}
	}
	if createDate != 0 {
		content += fmt.Sprintf(",\"create_date\": %d", createDate)
	} else {
		if beMasterSub {
			createDate = 7
			content += fmt.Sprintf(",\"create_date\": %d", createDate)
		}
	}
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAccountTransHisResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountTransHisResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetValidLeverRateAsync(data chan account.GetValidLeverRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetOrderLimitAsync(data chan account.GetOrderLimitResponse, orderPriceType string, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_limit", nil)

	// content
	content := fmt.Sprintf(",\"order_price_type\":\"%s\"", orderPriceType)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetOrderLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderLimitResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetFeeAsync(data chan account.GetFeeResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_fee", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetFeeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFeeResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetTransferLimitAsync(data chan account.GetTransferLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_transfer_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetTransferLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetTransferLimitResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetPositionLimitAsync(data chan account.GetPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_position_limit", nil)

	// content
	content := ""
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetPositionLimitResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetPositionLimitResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) GetApiTradingStatusAsync(data chan account.GetApiTradingStatusResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.GET_METHOD, "/linear-swap-api/v1/swap_api_trading_status", nil)

	// content is nil
	getResp, getErr := reqbuilder.HttpGet(url)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetApiTradingStatusResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetApiTradingStatusResponse error: %s", getErr)
	}
	data <- result
}
