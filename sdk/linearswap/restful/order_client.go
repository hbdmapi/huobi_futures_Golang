package restful

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	requestorder "huobi_futures_Golang/sdk/linearswap/restful/request/order"
	responseorder "huobi_futures_Golang/sdk/linearswap/restful/response/order"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
)

type OrderClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (oc *OrderClient) Init(accessKey string, secretKey string, host string) *OrderClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	oc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return oc
}

func (oc *OrderClient) IsolatedPlaceOrderAsync(data chan responseorder.PlaceOrderResponse, request requestorder.PlaceOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossPlaceOrderAsync(data chan responseorder.PlaceOrderResponse, request requestorder.PlaceOrderRequest) {
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order", nil)

	content, err := json.Marshal(request)
	if err != nil {
		log.Error("PlaceOrderRequest to json error: %v", err)
	}
	getResp, getErr := reqbuilder.HttpPost(url, string(content))
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedPlaceBatchOrderAsync(data chan responseorder.PlaceBatchOrderResponse, requests requestorder.BatchPlaceOrderRequest) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_batchorder", nil)

	// content
	bdata, err := json.Marshal(requests)
	if err != nil {
		log.Error("[] PlaceOrderRequest to json error: %v", err)
	}
	content := string(bdata)
	content = fmt.Sprintf("{\"orders_data\":%s}", content)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossPlaceBatchOrderAsync(data chan responseorder.PlaceBatchOrderResponse, requests requestorder.BatchPlaceOrderRequest) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_batchorder", nil)

	// content
	bdata, err := json.Marshal(requests)
	if err != nil {
		log.Error("[] PlaceOrderRequest to json error: %v", err)
	}
	content := string(bdata)
	content = fmt.Sprintf("{\"orders_data\":%s}", content)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.PlaceBatchOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to PlaceBatchOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedCancelOrderAsync(data chan responseorder.CancelOrderResponse, contractCode string,
	orderId string, clientOrderId string, offset string, direction string) {

	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cancel", nil)
	if orderId == "" && clientOrderId == "" {
		url = oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossCancelOrderAsync(data chan responseorder.CancelOrderResponse, contractCode string,
	orderId string, clientOrderId string, offset string, direction string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_cancel", nil)
	if orderId == "" && clientOrderId == "" {
		url = oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_cancelall", nil)
	}

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\": \"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\": \"%s\"", clientOrderId)
	}
	if offset != "" {
		content += fmt.Sprintf(",\"offset\": \"%s\"", offset)
	}
	if direction != "" {
		content += fmt.Sprintf(",\"direction\": \"%s\"", direction)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.CancelOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to CancelOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedSwitchLeverRateAsync(data chan responseorder.SwitchLeverRateResponse, contractCode string, leverRate int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_switch_lever_rate", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\", \"lever_rate\": %d}", contractCode, leverRate)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossSwitchLeverRateAsync(data chan responseorder.SwitchLeverRateResponse, contractCode string, leverRate int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_switch_lever_rate", nil)

	// content
	content := fmt.Sprintf("{\"contract_code\": \"%s\", \"lever_rate\": %d}", contractCode, leverRate)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.SwitchLeverRateResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to SwitchLeverRateResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetOrderInfoAsync(data chan responseorder.GetOrderInfoResponse, contractCode string,
	orderId string, clientOrderId string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_info", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\":\"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\":\"%s\"", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetOrderInfoAsync(data chan responseorder.GetOrderInfoResponse, contractCode string, orderId string, clientOrderId string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order_info", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\":\"%s\"", contractCode)
	if orderId != "" {
		content += fmt.Sprintf(",\"order_id\":\"%s\"", orderId)
	}
	if clientOrderId != "" {
		content += fmt.Sprintf(",\"client_order_id\":\"%s\"", clientOrderId)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetOrderInfoResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderInfoResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetOrderDetailAsync(data chan responseorder.GetOrderDetailResponse, contractCode string, orderId int64, createdAt int64,
	orderType int, pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_order_detail", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\", \"order_id\": %d", contractCode, orderId)
	if createdAt != 0 {
		content += fmt.Sprintf(",\"created_at\": %d", createdAt)
	}
	if orderType != 0 {
		content += fmt.Sprintf(",\"order_type\": %d", orderType)
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
	result := responseorder.GetOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetOrderDetailAsync(data chan responseorder.GetOrderDetailResponse, contractCode string, orderId int64, createdAt int64,
	orderType int, pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_order_detail", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\", \"order_id\": %d", contractCode, orderId)
	if createdAt != 0 {
		content += fmt.Sprintf(",\"created_at\": %d", createdAt)
	}
	if orderType != 0 {
		content += fmt.Sprintf(",\"order_type\": %d", orderType)
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
	result := responseorder.GetOrderDetailResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOrderDetailResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetOpenOrderAsync(data chan responseorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responseorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetOpenOrderAsync(data chan responseorder.GetOpenOrderResponse, contractCode string, pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_openorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\"", contractCode)
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
	result := responseorder.GetOpenOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetOpenOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisOrderAsync(data chan responseorder.GetHisOrderResponse, contractCode string, tradeType int, fcType int, status string,
	createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, fcType, status, createDate)
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
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisOrderAsync(data chan responseorder.GetHisOrderResponse, contractCode string, tradeType int, fcType int, status string,
	createDate int, pageIndex int, pageSize int, sortBy string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_hisorders", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": \"%d\",\"type\": \"%d\",\"status\": \"%s\",\"create_date\": %d", contractCode, tradeType, fcType, status, createDate)
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
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.GetHisOrderResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisOrderResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedGetHisMatchAsync(data chan responseorder.GetHisMatchResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_matchresults", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"create_date\": %d", contractCode, tradeType, createDate)
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
	result := responseorder.GetHisMatchResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossGetHisMatchAsync(data chan responseorder.GetHisMatchResponse, contractCode string, tradeType int, createDate int,
	pageIndex int, pageSize int) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_matchresults", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"trade_type\": %d,\"create_date\": %d", contractCode, tradeType, createDate)
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
	result := responseorder.GetHisMatchResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to GetHisMatchResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) IsolatedLightningCloseAsync(data chan responseorder.LightningCloseResponse, contractCode string, volume int, direction string,
	clientOrderId int64, orderPriceType string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_lightning_close_position", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"volume\": %d,\"direction\": \"%s\"", contractCode, volume, direction)
	if clientOrderId != 0 {
		content += fmt.Sprintf(",\"client_order_id\": %d", clientOrderId)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.LightningCloseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LightningCloseResponse error: %s", getErr)
	}
	data <- result
}

func (oc *OrderClient) CrossLightningCloseAsync(data chan responseorder.LightningCloseResponse, contractCode string, volume int, direction string,
	clientOrderId int64, orderPriceType string) {
	// url
	url := oc.PUrlBuilder.Build(linearswap.POST_METHOD, "/linear-swap-api/v1/swap_cross_lightning_close_position", nil)

	// content
	content := fmt.Sprintf(",\"contract_code\": \"%s\",\"volume\": %d,\"direction\": \"%s\"", contractCode, volume, direction)
	if clientOrderId != 0 {
		content += fmt.Sprintf(",\"client_order_id\": %d", clientOrderId)
	}
	if orderPriceType != "" {
		content += fmt.Sprintf(",\"order_price_type\": \"%s\"", orderPriceType)
	}
	if content != "" {
		content = fmt.Sprintf("{%s}", content[1:])
	}

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := responseorder.LightningCloseResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to LightningCloseResponse error: %s", getErr)
	}
	data <- result
}
