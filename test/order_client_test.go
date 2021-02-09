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

func TestOrderClient_IsolatedPlaceOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceOrderResponse)

	request := requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.45, 1, "sell", "open", 1, "limit", 0, 0, "", 0, 0, ""}
	go odClient.IsolatedPlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossPlaceOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceOrderResponse)

	request := requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.45, 1, "sell", "open", 1, "limit", 0.3, 0.3, "", 0.5, 0.5, ""}
	go odClient.CrossPlaceOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedPlaceBatchOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceBatchOrderResponse)

	request := requestorder.BatchPlaceOrderRequest{
		requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.45, 1, "buy", "open", 3, "limit", 0, 0, "", 0, 0, ""},
		requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.45, 1, "buy", "open", 3, "limit", 0, 0, "", 0, 0, ""},
	}
	go odClient.IsolatedPlaceBatchOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossPlaceBatchOrderAsync(t *testing.T) {
	data := make(chan responseorder.PlaceBatchOrderResponse)

	request := requestorder.BatchPlaceOrderRequest{
		requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.48, 1, "sell", "open", 3, "limit", 0.2, 0.2, "", 0.5, 0.5, ""},
		requestorder.PlaceOrderRequest{"XRP-USDT", 0, 0.48, 1, "sell", "open", 3, "limit", 0.1, 0.1, "", 0.3, 0.3, ""},
	}
	go odClient.CrossPlaceBatchOrderAsync(data, request)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedCancelOrderAsync(t *testing.T) {
	data := make(chan responseorder.CancelOrderResponse)

	go odClient.IsolatedCancelOrderAsync(data, "XRP-USDT", "779039570666164224", "", "open", "buy")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go odClient.IsolatedCancelOrderAsync(data, "XRP-USDT", "", "", "open", "buy")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossCancelOrderAsync(t *testing.T) {
	data := make(chan responseorder.CancelOrderResponse)

	go odClient.CrossCancelOrderAsync(data, "XRP-USDT", "779039570666164224", "", "open", "buy")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go odClient.CrossCancelOrderAsync(data, "XRP-USDT", "", "", "open", "buy")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedSwitchLeverRateAsync(t *testing.T) {
	data := make(chan responseorder.SwitchLeverRateResponse)

	go odClient.IsolatedSwitchLeverRateAsync(data, "XRP-USDT", 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossSwitchLeverRateAsync(t *testing.T) {
	data := make(chan responseorder.SwitchLeverRateResponse)

	go odClient.CrossSwitchLeverRateAsync(data, "XRP-USDT", 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetOrderInfoAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderInfoResponse)

	go odClient.IsolatedGetOrderInfoAsync(data, "XRP-USDT", "807205230369931264,807205439699243008", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetOrderInfoAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderInfoResponse)

	go odClient.CrossGetOrderInfoAsync(data, "XRP-USDT", "807205230369931264,807205439699243008", "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetOrderDetailAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderDetailResponse)

	go odClient.IsolatedGetOrderDetailAsync(data, "XRP-USDT", 807205230369931264, 0, 1, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetOrderDetailAsync(t *testing.T) {
	data := make(chan responseorder.GetOrderDetailResponse)

	go odClient.CrossGetOrderDetailAsync(data, "XRP-USDT", 807205439699243008, 0, 1, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetOpenOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetOpenOrderResponse)

	go odClient.IsolatedGetOpenOrderAsync(data, "XRP-USDT", 1, 10, "created_at", 2)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetOpenOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetOpenOrderResponse)

	go odClient.CrossGetOpenOrderAsync(data, "XRP-USDT", 1, 10, "created_at", 2)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetHisOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetHisOrderResponse)

	go odClient.IsolatedGetHisOrderAsync(data, "XRP-USDT", 0, 1, "0", 90, 1, 20, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetHisOrderAsync(t *testing.T) {
	data := make(chan responseorder.GetHisOrderResponse)

	go odClient.CrossGetHisOrderAsync(data, "XRP-USDT", 0, 1, "0", 90, 1, 20, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetHisOrderExactAsync(t *testing.T) {
	data := make(chan responseorder.GetHisOrderExactResponse)

	go odClient.IsolatedGetHisOrderExactAsync(data, "XRP-USDT", 0, 1, "0", "", 0, 0, 0, 0, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetHisOrderExactAsync(t *testing.T) {
	data := make(chan responseorder.GetHisOrderExactResponse)

	go odClient.CrossGetHisOrderExactAsync(data, "XRP-USDT", 0, 1, "0", "", 0, 0, 0, 0, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetHisMatchAsync(t *testing.T) {
	data := make(chan responseorder.GetHisMatchResponse)

	go odClient.IsolatedGetHisMatchAsync(data, "XRP-USDT", 0, 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetHisMatchAsync(t *testing.T) {
	data := make(chan responseorder.GetHisMatchResponse)

	go odClient.CrossGetHisMatchAsync(data, "XRP-USDT", 0, 1, 1, 20)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedGetHisMatchExactAsync(t *testing.T) {
	data := make(chan responseorder.GetHisMatchExactResponse)

	go odClient.IsolatedGetHisMatchExactAsync(data, "XRP-USDT", 0, 0, 0, 0, 0, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossGetHisMatchExactAsync(t *testing.T) {
	data := make(chan responseorder.GetHisMatchExactResponse)

	go odClient.CrossGetHisMatchExactAsync(data, "XRP-USDT", 0, 0, 0, 0, 0, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_IsolatedLightningCloseAsync(t *testing.T) {
	data := make(chan responseorder.LightningCloseResponse)

	go odClient.IsolatedLightningCloseAsync(data, "XRP-USDT", 1, "sell", 0, "lightning")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestOrderClient_CrossLightningCloseAsync(t *testing.T) {
	data := make(chan responseorder.LightningCloseResponse)

	go odClient.CrossLightningCloseAsync(data, "XRP-USDT", 1, "buy", 0, "lightning")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
