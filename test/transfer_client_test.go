package test

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"testing"
)

var tfClient restful.TransferClient

func init() {
	tfClient = restful.TransferClient{}
	tfClient.Init(config.AccessKey, config.SecretKey, "api.huobi.pro")
}

func TestTransferClient_TransferAsync(t *testing.T) {
	data := make(chan transfer.TransferResponse)

	go tfClient.TransferAsync(data, "spot", "linear-swap", 1, "BTC-USDT", "USDT")
	x, ok := <-data
	if !ok || x.Success != true {
		t.Logf("%d:%s", x.Code, x.Message)
		t.Fail()
	} else {
		t.Log(x)
	}
}
