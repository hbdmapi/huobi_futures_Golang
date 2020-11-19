package test

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	requestorder "huobi_futures_Golang/sdk/linearswap/restful/request/order"
	responseorder "huobi_futures_Golang/sdk/linearswap/restful/response/order"
	"testing"
)

var odClient restful.OrderClient

func init() {
	odClient = restful.OrderClient{}
	odClient.Init(config.AccessKey, config.SecretKey, config.Host)
}

func TestOrderClient_PlaceOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceOrderResponse)

	request := requestorder.PlaceOrderRequest{"EOS-USDT", 0, 2.1, 1, "buy", "open", 10, "limit"}
	go odClient.PlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_PlaceBatchOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceBatchOrderResponse)

	request := requestorder.BatchPlaceOrderRequest{
		requestorder.PlaceOrderRequest{"EOS-USDT", 0, 2.0, 1, "buy", "open", 10, "limit"},
		requestorder.PlaceOrderRequest{"EOS-USDT", 0, 2.1, 1, "buy", "open", 10, "limit"},
	}
	go odClient.PlaceBatchOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CancelOrderAsync(t *testing.T) {
	data := make(chan responseorder.CancelOrderResponse)

	go odClient.CancelOrderAsync(data, "EOS-USDT", "779039570666164224", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go odClient.CancelOrderAsync(data, "EOS-USDT", "", "")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_SwitchLeverRateAsync(t *testing.T) {
	data := make(chan responseorder.SwitchLeverRateResponse)

	go odClient.SwitchLeverRateAsync(data, "EOS-USDT", 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_GetOrderInfoAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderInfoResponse)

	go odClient.GetOrderInfoAsync(data, "EOS-USDT", "779039570666164224", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_GetOrderDetailAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderDetailResponse)

	go odClient.GetOrderDetailAsync(data, "EOS-USDT", 778777355463479296, 1605714381946, 1, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_GetOpenOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetOpenOrderResponse)

	go odClient.GetOpenOrderAsync(data, "EOS-USDT", 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_GetHisOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetHisOrderResponse)

	go odClient.GetHisOrderAsync(data, "EOS-USDT", 0, 1, "0", 5, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_GetHisMatchAsync(t *testing.T) {
	data := make(chan responseorder.GetHisMatchResponse)

	go odClient.GetHisMatchAsync(data, "EOS-USDT", 0, 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_LightningCloseAsync(t *testing.T) {
	data := make(chan responseorder.LightningCloseResponse)

	go odClient.LightningCloseAsync(data, "EOS-USDT", 1, "buy", 0, "lightning")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
