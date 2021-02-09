package test

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"testing"
)

var acClient restful.AccountClient

func init() {
	acClient = restful.AccountClient{}
	acClient.Init(config.AccessKey, config.SecretKey, config.Host)
}

func TestAccountClient_IsolatedGetAccountInfoAsync(t *testing.T) {

	data := make(chan account.GetAccountInfoResponse)

	go acClient.IsolatedGetAccountInfoAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.IsolatedGetAccountInfoAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetAccountInfoAsync(t *testing.T) {

	data := make(chan account.GetAccountInfoResponse)

	go acClient.CrossGetAccountInfoAsync(data, "USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.CrossGetAccountInfoAsync(data, "USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetAccountPositionAsync(t *testing.T) {

	data := make(chan account.GetAccountPositionResponse)

	go acClient.IsolatedGetAccountPositionAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.IsolatedGetAccountPositionAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetAccountPositionAsync(t *testing.T) {

	data := make(chan account.GetAccountPositionResponse)

	go acClient.CrossGetAccountPositionAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.CrossGetAccountPositionAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetAssetsPositionAsync(t *testing.T) {

	data := make(chan account.GetAssetsPositionResponse)

	go acClient.IsolatedGetAssetsPositionAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetAssetsPositionAsync(t *testing.T) {

	data := make(chan account.GetAssetsPositionResponseSingle)

	go acClient.CrossGetAssetsPositionAsync(data, "USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_SetSubAuthAsync(t *testing.T) {

	data := make(chan account.SetSubAuthResponse)

	go acClient.SetSubAuthAsync(data, "167891085", 1)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetSubAccountListAsync(t *testing.T) {

	data := make(chan account.GetSubAccountListResponse)

	go acClient.IsolatedGetSubAccountListResponseAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetAccountListAsync(t *testing.T) {

	data := make(chan account.GetSubAccountListResponse)

	go acClient.CrossGetSubAccountListAsync(data, "USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetSubAccountInfoListAsync(t *testing.T) {

	data := make(chan account.GetSubAccountInfoListResponse)

	go acClient.IsolatedGetSubAccountInfoListAsync(data, "BTC-USDT", 1, 100)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.CrossGetSubAccountInfoListAsync(data, "USDT", 1, 100)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_AccountTransferAsync(t *testing.T) {

	data := make(chan account.AccountTransferResponse)

	go acClient.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, config.SubUid, "master_to_sub")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, 0, "")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAccountTransHisAsync(t *testing.T) {

	data := make(chan account.GetAccountTransHisResponse)

	go acClient.GetAccountTransHisAsync(data, "BTC-USDT", false, "34", 10, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go acClient.GetAccountTransHisAsync(data, "BTC-USDT", true, "34", 10, 1, 10)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetFinancialRecordExactAsync(t *testing.T) {

	data := make(chan account.GetFinancialRecordExactResponse)

	go acClient.GetFinancialRecordExactAsync(data, "BTC-USDT", "BTC-USDT", "", 0, 0, 0, 0, "")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedSettlementRecordsAsync(t *testing.T) {

	data := make(chan account.IsolatedGetSettlementRecordsResponse)

	go acClient.IsolatedGetSettlementRecordsAsync(data, "ETH-USDT", 0, 0, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossSettlementRecordsAsync(t *testing.T) {

	data := make(chan account.CrossGetSettlementRecordsResponse)

	go acClient.CrossGetSettlementRecordsAsync(data, "USDT", 0, 0, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetValidLeverRateAsync(t *testing.T) {

	data := make(chan account.GetValidLeverRateResponse)

	go acClient.IsolatedGetValidLeverRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetValidLeverRateAsync(t *testing.T) {

	data := make(chan account.GetValidLeverRateResponse)

	go acClient.CrossGetValidLeverRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetOrderLimitAsync(t *testing.T) {

	data := make(chan account.GetOrderLimitResponse)

	go acClient.GetOrderLimitAsync(data, "limit", "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetFeeAsync(t *testing.T) {

	data := make(chan account.GetFeeResponse)

	go acClient.GetFeeAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetTransferLimitAsync(t *testing.T) {

	data := make(chan account.GetTransferLimitResponse)

	go acClient.IsolatedGetTransferLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_CrossGetTransferLimitAsync(t *testing.T) {

	data := make(chan account.GetTransferLimitResponse)

	go acClient.CrossGetTransferLimitAsync(data, "USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_IsolatedGetPositionLimitAsync(t *testing.T) {

	data := make(chan account.GetPositionLimitResponse)

	go acClient.IsolatedGetPositionLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetPositionLimitAsync(t *testing.T) {

	data := make(chan account.GetPositionLimitResponse)

	go acClient.CrossGetPositionLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetApiTradingStatusAsync(t *testing.T) {

	data := make(chan account.GetApiTradingStatusResponse)

	go acClient.GetApiTradingStatusAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
