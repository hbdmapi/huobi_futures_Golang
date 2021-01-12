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

func TestTriggerOrderClient_IsolatedPlaceOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.PlaceOrderResponse)

	request := requesttiggerorder.PlaceOrderRequest{"XRP-USDT", "le", 0.40, 0.40, "limit", 1, "buy", "open", 10}
	go todClient.IsolatedPlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_PlaceOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.PlaceOrderResponse)

	request := requesttiggerorder.PlaceOrderRequest{"XRP-USDT", "ge", 0.50, 0.50, "limit", 1, "sell", "open", 10}
	go todClient.CrossPlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_IsolatedCancelOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.CancelOrderResponse)

	go todClient.IsolatedCancelOrderAsync(data, "XRP-USDT", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_CrossCancelOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.CancelOrderResponse)

	go todClient.CrossCancelOrderAsync(data, "XRP-USDT", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_IsolatedGetOpenOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetOpenOrderResponse)

	go todClient.IsolatedGetOpenOrderAsync(data, "XRP-USDT", 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_CrossGetOpenOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetOpenOrderResponse)

	go todClient.CrossGetOpenOrderAsync(data, "XRP-USDT", 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_IsolatedGetHisOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetHisOrderResponse)

	go todClient.IsolatedGetHisOrderAsync(data, "XRP-USDT", 0, "0", 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_CrossGetHisOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetHisOrderResponse)

	go todClient.CrossGetHisOrderAsync(data, "XRP-USDT", 0, "0", 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_TpslOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.TpslOrderResponse)

	request := requesttiggerorder.TpslOrderRequest{"XRP-USDT", "buy", 1, 0.25, 0.25, "limit", 0.3, 0.3, "limit"}
	go odClient.IsolatedTpslOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	request = requesttiggerorder.TpslOrderRequest{"ETH-USDT", "buy", 1, 1000, 1000, "limit", 1100, 1100, "limit"}
	go odClient.CrossTpslOrderAsync(data, request)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
