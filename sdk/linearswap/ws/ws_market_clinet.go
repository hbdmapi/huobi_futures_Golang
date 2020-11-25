package ws

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/ws/response/market"
	"huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

type WSMarketClient struct {
	WebSocketOp
}

const default_id = "api"

func (wsMk *WSMarketClient) Init(host string) *WSMarketClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	wsMk.open("/linear-swap-ws", host, "", "", true)
	return wsMk
}

// -------------------------------------------------------------
// kline start

type OnSubKLineResponse func(*market.SubKLineResponse)
type OnReqKLineResponse func(*market.ReqKLineResponse)

func (wsMk *WSMarketClient) SubKLine(contractCode string, period string, callbackFun OnSubKLineResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubKLineResponse{}))
}

func (wsMk *WSMarketClient) UnsubKLine(contractCode string, period string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

func (wsMk *WSMarketClient) ReqKLine(contractCode string, period string, callbackFun OnReqKLineResponse, from int64, to int64, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	subData := wsbase.WSReqData{Req: ch, From: from, To: to, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.req(jdata, ch, callbackFun, reflect.TypeOf(market.ReqKLineResponse{}))
}

func (wsMk *WSMarketClient) UnreqKLine(contractCode string, period string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.kline.%s", contractCode, period)
	unsubData := wsbase.WSUnreqData{Unreq: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unreq(jdata, ch)
}

// kline end
//-------------------------------------------------------------

//-------------------------------------------------------------
// depth start

type OnSubDepthResponse func(*market.SubDepthResponse)

func (wsMk *WSMarketClient) SubDepth(contractCode string, fcType string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.depth.%s", contractCode, fcType)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubDepthResponse{}))
}

func (wsMk *WSMarketClient) UnsubDepth(contractCode string, fcType string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.depth.%s", contractCode, fcType)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

// depth end
//-------------------------------------------------------------

//-------------------------------------------------------------
// incrementa depth start

func (wsMk *WSMarketClient) SubIncrementalDepth(contractCode string, size string, callbackFun OnSubDepthResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.depth.size_%s.high_freq", contractCode, size)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubDepthResponse{}))
}

func (wsMk *WSMarketClient) UnsubIncrementalDepth(contractCode string, size string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.depth.size_%s.high_freq", contractCode, size)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

// incrementa depth end
//-------------------------------------------------------------

//-------------------------------------------------------------
// detail start
type OnSubDetailResponse func(*market.SubKLineResponse)

func (wsMk *WSMarketClient) SubDetail(contractCode string, callbackFun OnSubDetailResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.detail", contractCode)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubKLineResponse{}))
}

func (wsMk *WSMarketClient) UnsubDetail(contractCode string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.detail", contractCode)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

// detail end
//-------------------------------------------------------------

//-------------------------------------------------------------
// bbo start

type OnSubBBOResponse func(*market.SubBBOResponse)

func (wsMk *WSMarketClient) SubBBO(contractCode string, callbackFun OnSubBBOResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.bbo", contractCode)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubBBOResponse{}))
}

func (wsMk *WSMarketClient) UnsubBBO(contractCode string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.bbo", contractCode)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

// bbo end
//-------------------------------------------------------------

//-------------------------------------------------------------
// trade detail start

type OnSubTradeDetailResponse func(*market.SubTradeDetailResponse)
type OnReqTradeDetailResponse func(*market.ReqTradeDetailResponse)

func (wsMk *WSMarketClient) SubTradeDetail(contractCode string, callbackFun OnSubTradeDetailResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.sub(jdata, ch, callbackFun, reflect.TypeOf(market.SubTradeDetailResponse{}))
}

func (wsMk *WSMarketClient) UnsubTradeDetail(contractCode string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unsub(jdata, ch)
}

func (wsMk *WSMarketClient) ReqTradeDetail(contractCode string, callbackFun OnReqTradeDetailResponse, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	subData := wsbase.WSReqData{Req: ch, Id: id}
	jdata, _ := json.Marshal(subData)
	wsMk.req(jdata, ch, callbackFun, reflect.TypeOf(market.ReqTradeDetailResponse{}))
}

func (wsMk *WSMarketClient) UnreqTradeDetail(contractCode string, id string) {
	if id == "" {
		id = default_id
	}
	ch := fmt.Sprintf("market.%s.trade.detail", contractCode)
	unsubData := wsbase.WSUnreqData{Unreq: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)
	wsMk.unreq(jdata, ch)
}
