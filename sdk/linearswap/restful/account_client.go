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
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	ac.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return ac
}

func (ac *AccountClient) GetBalanceValuationAsync(data chan account.GetBalanceValuationResponse, valuationAsset string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_balance_valuation", nil)

	// content
	content := ""
	if valuationAsset != "" {
		content = fmt.Sprintf(",\"valuation_asset\": \"%s\"", valuationAsset)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetBalanceValuationResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetBalanceValuationResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAccountInfoAsync(data chan account.GetAccountInfoResponse, contractCode string, subUid int64) {
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
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAccountInfoAsync(data chan account.GetAccountInfoResponse, marginAccount string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info", nil)
	}

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
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
	result := account.GetAccountInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAccountInfoResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAccountPositionAsync(data chan account.GetAccountPositionResponse, contractCode string, subUid int64) {
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
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAccountPositionAsync(data chan account.GetAccountPositionResponse, contractCode string, subUid int64) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_info", nil)
	if subUid != 0 {
		url = ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_position_info", nil)
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
		log.Error("convert json to GetAccountPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetAssetsPositionAsync(data chan account.GetAssetsPositionResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\"}", contractCode)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAssetsPositionResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetAssetsPositionAsync(data chan account.GetAssetsPositionResponseSingle, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_account_position_info", nil)

	// content
	content := fmt.Sprintf("{\"margin_account\": \"%s\"}", marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetAssetsPositionResponseSingle{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAssetsPositionResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) SetSubAuthAsync(data chan account.SetSubAuthResponse, subUid string, subAuth int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_auth", nil)

	// content
	content := fmt.Sprintf("{\"sub_uid\": \"%s\", \"sub_auth\": %d}", subUid, subAuth)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.SetSubAuthResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SetSubAuthResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSubAccountListResponseAsync(data chan account.GetSubAccountListResponse, contractCode string) {
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
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetAllSubAssetsResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSubAccountListAsync(data chan account.GetSubAccountListResponse, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountListResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSubAccountInfoListAsync(data chan account.GetSubAccountInfoListResponse,
	contractCode string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_sub_account_info_list", nil)

	// content
	content := ""
	if contractCode != "" {
		content = fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSubAccountInfoListAsync(data chan account.GetSubAccountInfoListResponse,
	marginAccount string, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_sub_account_info_list", nil)

	// content
	content := ""
	if marginAccount != "" {
		content = fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	}
	if pageIndex != 0 {
		content = fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content = fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetSubAccountInfoListResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetSubAccountInfoListResponse error: %s", getErr)
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
		log.Error("convert json to AccountTransferResponse error: %s", jsonErr)
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
		log.Error("convert json to GetAccountTransHisResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) GetFinancialRecordExactAsync(data chan account.GetFinancialRecordExactResponse, marginAccount string,
	contractCode string, fcType string, startTime int64, endTime int64, fromId int64, size int, direct string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_financial_record_exact", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if contractCode != "" {
		content += fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	}
	if fcType != "" {
		content += fmt.Sprintf(",\"type\": \"%s\"", fcType)
	}
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
	}
	if fromId != 0 {
		content += fmt.Sprintf(",\"from_id\": %d", fromId)
	}
	if size != 0 {
		content += fmt.Sprintf(",\"size\": %d", size)
	}
	if direct != "" {
		content += fmt.Sprintf(",\"direct\": \"%s\"", direct)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetFinancialRecordExactResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetFinancialRecordExactResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetSettlementRecordsAsync(data chan account.IsolatedGetSettlementRecordsResponse, contractCode string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
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
	result := account.IsolatedGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to IsolatedGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetSettlementRecordsAsync(data chan account.CrossGetSettlementRecordsResponse, marginAccount string,
	startTime int64, endTime int64, pageIndex int, pageSize int) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_user_settlement_records", nil)

	// content
	content := fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
	if startTime != 0 {
		content += fmt.Sprintf(",\"start_time\": %d", startTime)
	}
	if endTime != 0 {
		content += fmt.Sprintf(",\"end_time\": %d", endTime)
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
	result := account.CrossGetSettlementRecordsResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CrossGetSettlementRecordsResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetValidLeverRateAsync(data chan account.GetValidLeverRateResponse, contractCode string) {
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
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetValidLeverRateAsync(data chan account.GetValidLeverRateResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_available_level_rate", nil)

	// content
	content := fmt.Sprintf("{ \"contract_code\": \"%s\" }", contractCode)
	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := account.GetValidLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetValidLeverRateResponse error: %s", jsonErr)
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
		log.Error("convert json to GetOrderLimitResponse error: %s", jsonErr)
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
		log.Error("convert json to GetFeeResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetTransferLimitAsync(data chan account.GetTransferLimitResponse, contractCode string) {
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
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetTransferLimitAsync(data chan account.GetTransferLimitResponse, marginAccount string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_transfer_limit", nil)

	// content
	content := ""
	if marginAccount != "" {
		content += fmt.Sprintf(",\"margin_account\": \"%s\"", marginAccount)
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
		log.Error("convert json to GetTransferLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) IsolatedGetPositionLimitAsync(data chan account.GetPositionLimitResponse, contractCode string) {
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
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
	}
	data <- result
}

func (ac *AccountClient) CrossGetPositionLimitAsync(data chan account.GetPositionLimitResponse, contractCode string) {
	// ulr
	url := ac.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_position_limit", nil)

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
		log.Error("convert json to GetPositionLimitResponse error: %s", jsonErr)
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
		log.Error("convert json to GetApiTradingStatusResponse error: %s", jsonErr)
	}
	data <- result
}
