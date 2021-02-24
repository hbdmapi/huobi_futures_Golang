package test

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/market"
	"testing"
)

var mkClient restful.MarketClient

func init() {
	mkClient = restful.MarketClient{}
	mkClient.Init(config.Host)
}

func TestMarketClient_GetContractInfoAsync(t *testing.T) {

	data := make(chan market.GetContractInfoResponse)

	go mkClient.GetContractInfoAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetContractInfoAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetIndexAsync(t *testing.T) {
	data := make(chan market.GetIndexResponse)

	go mkClient.GetIndexAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetIndexAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetPriceLimitAsync(t *testing.T) {
	data := make(chan market.GetPriceLimitResponse)

	go mkClient.GetPriceLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetOpenInterestAsync(t *testing.T) {
	data := make(chan market.GetOpenInterestResponse)

	go mkClient.GetOpenInterestAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetOpenInterestAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetDepthAsync(t *testing.T) {
	data := make(chan market.GetDepthResponse)

	go mkClient.GetDepthAsync(data, "BTC-USDT", "step0")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetKLineAsync(t *testing.T) {
	data := make(chan market.GetKLineResponse)

	go mkClient.GetKLineAsync(data, "BTC-USDT", "1min", 10, 0, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetKLineAsync(data, "BTC-USDT", "1min", 0, 1605603881, 1605683881)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetMarkPriceKLineAsync(t *testing.T) {
	data := make(chan market.GetStrKLineResponse)

	go mkClient.GetMarkPriceKLineAsync(data, "BTC-USDT", "1min", 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetMergedAsync(t *testing.T) {
	data := make(chan market.GetMergedResponse)

	go mkClient.GetMergedAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetBatchMergedAsync(t *testing.T) {
	data := make(chan market.GetBatchMergedResponse)

	go mkClient.GetBatchMergedAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetTradeAsync(t *testing.T) {
	data := make(chan market.GetTradeResponse)

	go mkClient.GetTradeAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetHisTradeAsync(t *testing.T) {
	data := make(chan market.GetHisTradeResponse)

	go mkClient.GetHisTradeAsync(data, "BTC-USDT", 2)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetRiskInfoAsync(t *testing.T) {
	data := make(chan market.GetRiskInfoResponse)

	go mkClient.GetRiskInfoAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetRiskInfoAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetInsuranceFundAsync(t *testing.T) {
	data := make(chan market.GetInsuranceFundResponse)

	go mkClient.GetInsuranceFundAsync(data, "BTC-USDT", 0, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetInsuranceFundAsync(data, "BTC-USDT", 1, 2)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_IsolatedGetAdjustFactorFundAsync(t *testing.T) {
	data := make(chan market.GetAdjustFactorFundResponse)

	go mkClient.IsolatedGetAdjustFactorFundAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.IsolatedGetAdjustFactorFundAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_CrossGetAdjustFactorFundAsync(t *testing.T) {
	data := make(chan market.GetAdjustFactorFundResponse)

	go mkClient.CrossGetAdjustFactorFundAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.CrossGetAdjustFactorFundAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetHisOpenInterestAsync(t *testing.T) {
	data := make(chan market.GetHisOpenInterestResponse)

	go mkClient.GetHisOpenInterestAsync(data, "BTC-USDT", "60min", 1, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetHisOpenInterestAsync(data, "BTC-USDT", "60min", 2, 10)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetLadderMarginAsync(t *testing.T) {
	data := make(chan market.GetLadderMarginResponse)

	go mkClient.IsolatedGetLadderMarginAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.CrossGetLadderMarginAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetEliteAccountRatioAsync(t *testing.T) {
	data := make(chan market.GetEliteRatioResponse)

	go mkClient.GetEliteAccountRatioAsync(data, "BTC-USDT", "15min")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetElitePositionRatioAsync(t *testing.T) {
	data := make(chan market.GetEliteRatioResponse)

	go mkClient.GetElitePositionRatioAsync(data, "BTC-USDT", "15min")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_IsolatedGetApiStatusAsync(t *testing.T) {
	data := make(chan market.GetApiStateResponse)

	go mkClient.IsolatedGetApiStateAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.IsolatedGetApiStateAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_CrossGetTransferStateAsync(t *testing.T) {
	data := make(chan market.GetTransferStateResponse)

	go mkClient.CrossGetTransferStateAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_CrossTradeStateAsync(t *testing.T) {
	data := make(chan market.GetTradeStateResponse)

	go mkClient.CrossGetTradeStateAsync(data, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.CrossGetTradeStateAsync(data, "BTC-USDT")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetFundingRateAsync(t *testing.T) {
	data := make(chan market.GetFundingRateResponse)

	go mkClient.GetFundingRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetBatchFundingRateAsync(t *testing.T) {
	data := make(chan market.GetBatchFundingRateResponse)

	go mkClient.GetBatchFundingRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetHisFundingRateAsync(t *testing.T) {
	data := make(chan market.GetHisFundingRateResponse)

	go mkClient.GetHisFundingRateAsync(data, "BTC-USDT", 0, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetHisFundingRateAsync(data, "BTC-USDT", 2, 6)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetLiquidationOrdersAsync(t *testing.T) {
	data := make(chan market.GetLiquidationOrdersResponse)

	go mkClient.GetLiquidationOrdersAsync(data, "BTC-USDT", 0, 7, 0, 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetLiquidationOrdersAsync(data, "BTC-USDT", 0, 7, 2, 6)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetPremiumIndexKLineAsync(t *testing.T) {
	data := make(chan market.GetStrKLineResponse)

	go mkClient.GetPremiumIndexKLineAsync(data, "BTC-USDT", "5min", 2)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetEstimatedRateKLineAsync(t *testing.T) {
	data := make(chan market.GetStrKLineResponse)

	go mkClient.GetEstimatedRateKLineAsync(data, "BTC-USDT", "5min", 2)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetBasisAsync(t *testing.T) {
	data := make(chan market.GetBasisResponse)

	go mkClient.GetBasisAsync(data, "BTC-USDT", "5min", 1, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go mkClient.GetBasisAsync(data, "BTC-USDT", "5min", 1, "close")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestMarketClient_GetEstimatedSettlementPriceAsync(t *testing.T) {
	data := make(chan market.GetEstimatedSettlementPriceResponse)

	go mkClient.GetEstimatedSettlementPriceAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
