package test

import (
	"huobi_futures_Golang/sdk/linearswap/ws"
	"huobi_futures_Golang/sdk/linearswap/ws/response/index"
	"testing"
	"time"
)

var wsixClient *ws.WSIndexClient

func init() {
	//wsixClient = new(ws.WSIndexClient).Init("")
}

func TestWSIndexClient_SubPremiumIndexKLine(t *testing.T) {
	wsixClient.SubPremiumIndexKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnsubPremiumIndexKLine("BTC-USDT", "15min", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSIndexClient_ReqPremiumIndexKLine(t *testing.T) {
	wsixClient.ReqPremiumIndexKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		t.Log(*m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnreqPremiumIndexKLine("BTC-USDT", "15min", 1604395758, 1604396758, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSIndexClient_SubEstimatedRateKLine(t *testing.T) {
	wsixClient.SubEstimatedRateKLine("BTC-USDT", "15min", func(m *index.SubIndexKLineResponse) {
		t.Log(*m)
	}, "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnsubEstimatedRateKLine("BTC-USDT", "15min", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSIndexClient_ReqEstimatedRateKLine(t *testing.T) {
	wsixClient.ReqEstimatedRateKLine("BTC-USDT", "15min", func(m *index.ReqIndexKLineResponse) {
		t.Log(*m)
	}, 1604395758, 1604396758, "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnreqEstimatedRateKLine("BTC-USDT", "15min", 1604395758, 1604396758, "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSIndexClient_SubBasis(t *testing.T) {
	wsixClient.SubBasis("BTC-USDT", "15min", func(m *index.SubBasiesResponse) {
		t.Log(*m)
	}, "", "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnsubBasis("BTC-USDT", "15min", "", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TestWSIndexClient_ReqBasis(t *testing.T) {
	wsixClient.ReqBasis("BTC-USDT", "15min", func(m *index.ReqBasisResponse) {
		t.Log(*m)
	}, 1604395758, 1604396758, "", "")
	time.Sleep(time.Duration(20) * time.Second)
	wsixClient.UnreqBasis("BTC-USDT", "15min", 1604395758, 1604396758, "", "")
	time.Sleep(time.Duration(10) * time.Second)
}
