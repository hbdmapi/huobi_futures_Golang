package test

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	requesttiggerorder "huobi_futures_Golang/sdk/linearswap/restful/request/triggerorder"
	responsetriggerorder "huobi_futures_Golang/sdk/linearswap/restful/response/triggerorder"
	"testing"
)

var todClient restful.TriggerOrderClient

func init() {
	todClient = restful.TriggerOrderClient{}
	todClient.Init(config.AccessKey, config.SecretKey, config.Host)
}

func TestTriggerOrderClient_PlaceOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.PlaceOrderResponse)

	request := requesttiggerorder.PlaceOrderRequest{"ETH-USDT", "le", 450, 450, "limit", 1, "buy", "open", 20}
	go todClient.PlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_CancelOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.CancelOrderResponse)

	go todClient.CancelOrderAsync(data, "ETH-USDT", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_GetOpenOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetOpenOrderResponse)

	go todClient.GetOpenOrderAsync(data, "ETH-USDT", 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_GetHisOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetHisOrderResponse)

	go todClient.GetHisOrderAsync(data, "ETH-USDT", 0, "0", 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
