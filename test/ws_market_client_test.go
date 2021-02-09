package test

import (
	"huobi_futures_Golang/sdk/linearswap/ws"
	"huobi_futures_Golang/sdk/linearswap/ws/response/market"
	"testing"
	"time"
)

var wsmkClient *ws.WSMarketClient

func init() {
	wsmkClient = new(ws.WSMarketClient).Init("api.hbdm.com")
}

func TestWSMarketClient_SubKLine(t *testing.T) {
	wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSMarketClient_ReqKLine(t *testing.T) {
	wsmkClient.ReqKLine("BTC-USDT", "1min", func(m *market.ReqKLineResponse) {
		t.Log(*m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSMarketClient_SubDepth(t *testing.T) {
	wsmkClient.SubDepth("BTC-USDT", "step0", func(m *market.SubDepthResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func TestWSMarketClient_SubIncrementalDepth(t *testing.T) {
	wsmkClient.SubIncrementalDepth("BTC-USDT", "20", func(m *market.SubDepthResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func TestWSMarketClient_SubDetail(t *testing.T) {
	wsmkClient.SubDetail("BTC-USDT", func(m *market.SubKLineResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func TestWSMarketClient_SubBBO(t *testing.T) {
	wsmkClient.SubBBO("BTC-USDT", func(m *market.SubBBOResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func TestWSMarketClient_SubTradeDetail(t *testing.T) {
	wsmkClient.SubTradeDetail("BTC-USDT", func(m *market.SubTradeDetailResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}

func TestWSMarketClient_ReqTradeDetail(t *testing.T) {
	wsmkClient.ReqTradeDetail("BTC-USDT", func(m *market.ReqTradeDetailResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(5) * time.Second)
}
