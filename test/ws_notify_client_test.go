package test

import (
	"huobi_futures_Golang/sdk/linearswap/ws"
	"huobi_futures_Golang/sdk/linearswap/ws/response/notify"
	"testing"
	"time"
)

var wsnfClient *ws.WSNotifyClient

func init() {
	wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
}

func TestWSNotifyClient_SubOrders(t *testing.T) {
	wsnfClient.IsolatedSubOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	wsnfClient.CrossSubOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.IsolatedUnsubOrders("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubAcounts(t *testing.T) {
	wsnfClient.IsolatedSubAcounts("*", func(m *notify.SubAccountsResponse) {
		t.Log(*m)
	}, "")
	wsnfClient.CrossSubAcounts("*", func(m *notify.SubAccountsResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.IsolatedUnsubAccounts("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubPositions(t *testing.T) {
	wsnfClient.IsolatedSubPositions("xrp-usdt", func(m *notify.SubPositionsResponse) {
		t.Log(*m)
	}, "")
	wsnfClient.CrossSubPositions("xrp-usdt", func(m *notify.SubPositionsResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.IsolatdUnsubPositions("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubMatchOrders(t *testing.T) {
	wsnfClient.IsolatedSubMatchOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	wsnfClient.CrossSubMatchOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.IsolatedUnsubMathOrders("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubLiquidationOrders(t *testing.T) {
	wsnfClient.SubLiquidationOrders("*", func(m *notify.SubLiquidationOrdersResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.UnsubLiquidationOrders("*", "")
	time.Sleep(time.Duration(500) * time.Second)
}

func TestWSNotifyClient_SubFundingRate(t *testing.T) {
	wsnfClient.SubFundingRate("BTC-USDT", func(m *notify.SubFundingRateResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
	wsnfClient.UnsubFundingRate("BTC-USDT", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSNotifyClient_SubContractInfo(t *testing.T) {
	wsnfClient.SubContractInfo("*", func(m *notify.SubContractInfoResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
	wsnfClient.UnsubContractInfo("*", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSNotifyClient_SubTriggerOrder(t *testing.T) {
	wsnfClient.IsolatedSubTriggerOrder("*", func(m *notify.SubTriggerOrderResponse) {
		t.Log(*m)
	}, "")
	wsnfClient.CrossSubTriggerOrder("*", func(m *notify.SubTriggerOrderResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(100) * time.Second)
	wsnfClient.IsolatedUnsubTriggerOrder("*", "")
	time.Sleep(time.Duration(100) * time.Second)
}
