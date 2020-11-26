package ws

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/ws/response/notify"
	"huobi_futures_Golang/sdk/wsbase"
	"reflect"
	"strings"
)

type WSNotifyClient struct {
	WebSocketOp
}

func (wsNf *WSNotifyClient) Init(accessKey string, secretKey string, host string) *WSNotifyClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	wsNf.open("/linear-swap-notification", host, accessKey, secretKey, true)
	return wsNf
}

// -------------------------------------------------------------
// orders start

type OnSubOrdersResponse func(*notify.SubOrdersResponse)

func (wsNf *WSNotifyClient) SubOrders(contractCode string, callbackFun OnSubOrdersResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}
	ch := fmt.Sprintf("orders.%s", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubOrdersResponse{}))
}

func (wsNf *WSNotifyClient) UnsubOrders(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("orders.%s", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// accounts start

type OnSubAccountsResponse func(*notify.SubAccountsResponse)

func (wsNf *WSNotifyClient) SubAcounts(contractCode string, callbackFun OnSubAccountsResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubAccountsResponse{}))
}

func (wsNf *WSNotifyClient) UnsubAccounts(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("accounts.%s", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// accounts end
//-------------------------------------------------------------

// -------------------------------------------------------------
// positions start

type OnSubPositionsResponse func(*notify.SubPositionsResponse)

func (wsNf *WSNotifyClient) SubPositions(contractCode string, callbackFun OnSubPositionsResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubPositionsResponse{}))
}

func (wsNf *WSNotifyClient) UnsubPositions(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("positions.%s", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// positions end
//-------------------------------------------------------------

// -------------------------------------------------------------
// match orders start

type OnSubMatchOrdersResponse func(*notify.SubOrdersResponse)

func (wsNf *WSNotifyClient) SubMatchOrders(contractCode string, callbackFun OnSubMatchOrdersResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders.%s", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubOrdersResponse{}))
}

func (wsNf *WSNotifyClient) UnsubMathOrders(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	contractCode = strings.ToLower(contractCode)
	ch := fmt.Sprintf("matchOrders.%s", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// match orders end
//-------------------------------------------------------------

// -------------------------------------------------------------
// liquidation orders start

type OnSubLiquidationOrdersResponse func(*notify.SubLiquidationOrdersResponse)

func (wsNf *WSNotifyClient) SubLiquidationOrders(contractCode string, callbackFun OnSubLiquidationOrdersResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubLiquidationOrdersResponse{}))
}

func (wsNf *WSNotifyClient) UnsubLiquidationOrders(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.liquidation_orders", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// liquidation orders end
//-------------------------------------------------------------

// -------------------------------------------------------------
// funding rate start

type OnSubFundingRateResponse func(*notify.SubFundingRateResponse)

func (wsNf *WSNotifyClient) SubFundingRate(contractCode string, callbackFun OnSubFundingRateResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.funding_rate", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubFundingRateResponse{}))
}

func (wsNf *WSNotifyClient) UnsubFundingRate(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.funding_rate", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// funding rate end
//-------------------------------------------------------------

// -------------------------------------------------------------
// contract info start

type OnSubContractInfoResponse func(*notify.SubContractInfoResponse)

func (wsNf *WSNotifyClient) SubContractInfo(contractCode string, callbackFun OnSubContractInfoResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubContractInfoResponse{}))
}

func (wsNf *WSNotifyClient) UnsubContractInfo(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.contract_info", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// fcontract info end
//-------------------------------------------------------------

// -------------------------------------------------------------
// trigger order start

type OnSubTriggerOrderResponse func(*notify.SubTriggerOrderResponse)

func (wsNf *WSNotifyClient) SubTriggerOrder(contractCode string, callbackFun OnSubTriggerOrderResponse, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", contractCode)
	opData := wsbase.WSOpData{Op: "sub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(notify.SubTriggerOrderResponse{}))
}

func (wsNf *WSNotifyClient) UnsubTriggerOrder(contractCode string, cid string) {
	if cid == "" {
		cid = linearswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("trigger_order.%s", contractCode)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}
