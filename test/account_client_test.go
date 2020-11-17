package test

import (
	"encoding/json"
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/account"
	"os"
	"testing"
)

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	AccountId int64
	SubUid    int64
}

func GetConfig() *Config {

	filePtr, err := os.Open("config.json")
	if err != nil {
		return nil
	}
	defer filePtr.Close()

	var data Config

	// 创建json解码器
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&data)

	return &data
}

var client restful.AccountClient
var config *Config

func Init() {
	client = restful.AccountClient{}
	config = GetConfig()
	client.Init(config.AccessKey, config.SecretKey, config.Host)
}

func TestAccountClient_GetAccountAssetsAsync(t *testing.T) {
	Init()

	data := make(chan account.GetAccountAssetsResponse)

	go client.GetAccountAssetsAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go client.GetAccountAssetsAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAccountPositionAsync(t *testing.T) {
	Init()

	data := make(chan account.GetAccountPositionResponse)

	go client.GetAccountPositionAsync(data, "BTC-USDT", 0)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go client.GetAccountPositionAsync(data, "BTC-USDT", config.SubUid)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAllSubAssetsAsync(t *testing.T) {
	Init()

	data := make(chan account.GetAllSubAssetsResponse)

	go client.GetAllSubAssetsAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_AccountTransferAsync(t *testing.T) {
	Init()

	data := make(chan account.AccountTransferResponse)

	go client.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, config.SubUid, "master_to_sub")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go client.AccountTransferAsync(data, "USDT", "BTC-USDT", "ETH-USDT", 1, 0, "")
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetAccountTransHisAsync(t *testing.T) {
	Init()

	data := make(chan account.GetAccountTransHisResponse)

	go client.GetAccountTransHisAsync(data, "BTC-USDT", false, "34", 10, 1, 10)
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}

	go client.GetAccountTransHisAsync(data, "BTC-USDT", true, "34", 10, 1, 10)
	x, ok = <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetValidLeverRateAsync(t *testing.T) {
	Init()

	data := make(chan account.GetValidLeverRateResponse)

	go client.GetValidLeverRateAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetOrderLimitAsync(t *testing.T) {
	Init()

	data := make(chan account.GetOrderLimitResponse)

	go client.GetOrderLimitAsync(data, "limit", "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetFeeAsync(t *testing.T) {
	Init()

	data := make(chan account.GetFeeResponse)

	go client.GetFeeAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetTransferLimitAsync(t *testing.T) {
	Init()

	data := make(chan account.GetTransferLimitResponse)

	go client.GetTransferLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetPositionLimitAsync(t *testing.T) {
	Init()

	data := make(chan account.GetPositionLimitResponse)

	go client.GetPositionLimitAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}

func TestAccountClient_GetApiTradingStatusAsync(t *testing.T) {
	Init()

	data := make(chan account.GetApiTradingStatusResponse)

	go client.GetApiTradingStatusAsync(data, "BTC-USDT")
	x, ok := <-data
	if !ok || x.Status != "ok" {
		t.Logf("%d:%s", x.ErrorCode, x.ErrorMessage)
		t.Fail()
	} else {
		t.Log(x)
	}
}
