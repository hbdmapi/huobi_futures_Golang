package restful

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
)

type TransferClient struct {
	PUrlBuilder *reqbuilder.PrivateUrlBuilder
}

func (tc *TransferClient) Init(accessKey string, secretKey string, host string) *TransferClient {
	if host == "" {
		host = linearswap.SPOT_DEFAULT_HOST
	}
	tc.PUrlBuilder = new(reqbuilder.PrivateUrlBuilder).Init(accessKey, secretKey, host)
	return tc
}

func (tc *TransferClient) TransferAsync(data chan transfer.TransferResponse, from string, to string, amount float32, marginAccount string, currency string) {
	if currency == "" {
		currency = "USDT"
	}

	// ulr
	url := tc.PUrlBuilder.Build(linearswap.POST_METHOD, "/v2/account/transfer", nil)

	// content
	content := fmt.Sprintf("{ \"from\":\"%s\", \"to\":\"%s\", \"currency\":\"%s\", \"amount\":%f, \"margin-account\":\"%s\" }", from, to, currency, amount, marginAccount)

	getResp, getErr := reqbuilder.HttpPost(url, content)
	if getErr != nil {
		log.Error("http get error: %s", getErr)
	}
	result := transfer.TransferResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		log.Error("convert json to TransferResponse error: %s", getErr)
	}
	data <- result
}
