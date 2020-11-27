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
	wsnfClient.SubOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.UnsubOrders("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubAcounts(t *testing.T) {
	wsnfClient.SubAcounts("*", func(m *notify.SubAccountsResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.UnsubAccounts("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubPositions(t *testing.T) {
	wsnfClient.SubPositions("*", func(m *notify.SubPositionsResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.UnsubPositions("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func TestWSNotifyClient_SubMatchOrders(t *testing.T) {
	wsnfClient.SubMatchOrders("*", func(m *notify.SubOrdersResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.UnsubMathOrders("*", "")
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
	wsnfClient.SubTriggerOrder("*", func(m *notify.SubTriggerOrderResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(100) * time.Second)
	wsnfClient.UnsubTriggerOrder("*", "")
	time.Sleep(time.Duration(100) * time.Second)
}
