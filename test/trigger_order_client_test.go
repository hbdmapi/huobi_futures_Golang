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

	request := requesttiggerorder.TpslOrderRequest{"ADA-USDT", "buy", 1, 0.25, 0.25, "limit", 0.31, 0.31, "limit"}
	go odClient.IsolatedTpslOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	request = requesttiggerorder.TpslOrderRequest{"ADA-USDT", "buy", 1, 0.25, 0.25, "limit", 0.31, 0.31, "limit"}
	go odClient.CrossTpslOrderAsync(data, request)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_TpslCancelAsync(t *testing.T) {
	data := make(chan responsetriggerorder.CancelOrderResponse)

	go todClient.IsolatedTpslCancelAsync(data, "ADA-USDT", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go todClient.CrossTpslCancelAsync(data, "ADA-USDT", "")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_GetTpslOpenOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetOpenOrderResponse)

	go todClient.IsolatedGetTpslOpenOrderAsync(data, "ADA-USDT", 0, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go todClient.CrossGetTpslOpenOrderAsync(data, "ADA-USDT", 0, 0)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_GetTpslHisOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetHisOrderResponse)

	go todClient.IsolatedGetTpslHisOrderAsync(data, "ADA-USDT", "0", 10, 0, 0, "created_at")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go todClient.CrossGetTpslHisOrderAsync(data, "ADA-USDT", "0", 10, 0, 0, "created_at")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestTriggerOrderClient_GetRelationTpslOrderAsync(t *testing.T) {
	data := make(chan responsetriggerorder.GetRelationTpslOrderResponse)

	go todClient.IsolatedGetRelationTpslOrderAsync(data, "ADA-USDT", 798613002479038466)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go todClient.CrossGetRelationTpslOrderAsync(data, "ADA-USDT", 798613002479038466)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
