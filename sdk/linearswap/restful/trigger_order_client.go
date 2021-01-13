package restful

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	requesttiggerorder "huobi_futures_Golang/sdk/linearswap/restful/request/triggerorder"
	responsetriggerorder "huobi_futures_Golang/sdk/linearswap/restful/response/triggerorder"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
)

type TriggerOrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (toc *TriggerOrderClient) Init(accessKey string, secretKey string, host string) *TriggerOrderClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	toc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return toc
}

func (toc *TriggerOrderClient) IsolatedPlaceOrderAsync(data chan responsetriggerorder.PlaceOrderResponse, request requesttiggerorder.PlaceOrderRequest) {
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossPlaceOrderAsync(data chan responsetriggerorder.PlaceOrderResponse, request requesttiggerorder.PlaceOrderRequest) {
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedCancelOrderAsync(data chan responsetriggerorder.CancelOrderResponse, contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossCancelOrderAsync(data chan responsetriggerorder.CancelOrderResponse, contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetOpenOrderAsync(data chan responsetriggerorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetOpenOrderAsync(data chan responsetriggerorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetHisOrderAsync(data chan responsetriggerorder.GetHisOrderResponse, contractCode string, tradeType int, status string, createDate int,
	pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_trigger_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetHisOrderAsync(data chan responsetriggerorder.GetHisOrderResponse, contractCode string, tradeType int, status string, createDate int,
	pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_trigger_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\": \"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedTpslOrderAsync(data chan responsetriggerorder.TpslOrderResponse, request requesttiggerorder.TpslOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.TpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to TpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossTpslOrderAsync(data chan responsetriggerorder.TpslOrderResponse, request requesttiggerorder.TpslOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.TpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to TpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedTpslCancelAsync(data chan responsetriggerorder.CancelOrderResponse, contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossTpslCancelAsync(data chan responsetriggerorder.CancelOrderResponse, contractCode string, orderId string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_cancel", nil)
	if orderId == "" {
		url = toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetTpslOpenOrderAsync(data chan responsetriggerorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetTpslOpenOrderAsync(data chan responsetriggerorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetTpslHisOrderAsync(data chan responsetriggerorder.GetHisOrderResponse,
	contractCode string, status string, createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_tpsl_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"status\":\"%s\", \"create_date\":%d", contractCode, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\":\"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetTpslHisOrderAsync(data chan responsetriggerorder.GetHisOrderResponse,
	contractCode string, status string, createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_tpsl_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"status\":\"%s\", \"create_date\":%d", contractCode, status, createDate)
	if pageIndex != 0 {
		content += fmt.Sprintf(",\"page_index\": %d", pageIndex)
	}
	if pageSize != 0 {
		content += fmt.Sprintf(",\"page_size\": %d", pageSize)
	}
	if sortBy != "" {
		content += fmt.Sprintf(",\"sort_by\":\"%s\"", sortBy)
	}
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) IsolatedGetRelationTpslOrderAsync(data chan responsetriggerorder.GetRelationTpslOrderResponse,
	contractCode string, orderId int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_relation_tpsl_order", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"order_id\":%d", contractCode, orderId)
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}

func (toc *TriggerOrderClient) CrossGetRelationTpslOrderAsync(data chan responsetriggerorder.GetRelationTpslOrderResponse,
	contractCode string, orderId int) {
	// url
	url := toc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_relation_tpsl_order", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\", \"order_id\":%d", contractCode, orderId)
	if content != "" {
		content = fmt.Sprintf("{ %s }", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responsetriggerorder.GetRelationTpslOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetRelationTpslOrderResponse error: %s", getErr)
	}
	data <- result
}
