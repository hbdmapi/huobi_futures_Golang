package ws

import (
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/linearswap/ws/response/index"
	"huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

type WSIndexClient struct {
	WebSocketOp
}

func (wsIx *WSIndexClient) Init(host string) *WSIndexClient {
	if host == "" {
		host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}
	wsIx.open("/ws_index", host, "", "", true)
	return wsIx
}

// -------------------------------------------------------------
// premium index kline start
type OnSubPremiumIndexKLineResponse func(*index.SubIndexKLineResponse)
type OnReqPremiumIndexKLineResponse func(*index.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubPremiumIndexKLine(contractCode string, period string, callbackFun OnSubPremiumIndexKLineResponse, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) UnsubPremiumIndexKLine(contractCode string, period string, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)

	wsIx.unsub(jdata, ch)
}

func (wsIx *WSIndexClient) ReqPremiumIndexKLine(contractCode string, period string, callbackFun OnReqPremiumIndexKLineResponse,
	from int64, to int64, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) UnreqPremiumIndexKLine(contractCode string, period string, from int64, to int64, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.premium_index.%s", contractCode, period)
	unreqData := wsbase.WSUnreqData{Unreq: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(unreqData)

	wsIx.unreq(jdata, ch)
}

// premium index kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// estimated rate kline start

type OnSubEstimatedRateResponse func(*index.SubIndexKLineResponse)
type OnReqEstimatedRateResponse func(*index.ReqIndexKLineResponse)

func (wsIx *WSIndexClient) SubEstimatedRateKLine(contractCode string, period string, callbackFun OnSubEstimatedRateResponse, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) UnsubEstimatedRateKLine(contractCode string, period string, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)

	wsIx.unsub(jdata, ch)
}

func (wsIx *WSIndexClient) ReqEstimatedRateKLine(contractCode string, period string, callbackFun OnReqEstimatedRateResponse,
	from int64, to int64, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqIndexKLineResponse{}))
}

func (wsIx *WSIndexClient) UnreqEstimatedRateKLine(contractCode string, period string, from int64, to int64, id string) {
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.estimated_rate.%s", contractCode, period)
	unreqData := wsbase.WSUnreqData{Unreq: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(unreqData)

	wsIx.unreq(jdata, ch)
}

// estimated rate kline end
//-------------------------------------------------------------

// -------------------------------------------------------------
// basis start
type OnSubBasisResponse func(*index.SubBasiesResponse)
type OnReqBasisResponse func(*index.ReqBasisResponse)

func (wsIx *WSIndexClient) SubBasis(contractCode string, period string, callbackFun OnSubBasisResponse, basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	subData := wsbase.WSSubData{Sub: ch, Id: id}
	jdata, _ := json.Marshal(subData)

	wsIx.sub(jdata, ch, callbackFun, reflect.TypeOf(index.SubBasiesResponse{}))
}

func (wsIx *WSIndexClient) UnsubBasis(contractCode string, period string, basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	unsubData := wsbase.WSUnsubData{Unsub: ch, Id: id}
	jdata, _ := json.Marshal(unsubData)

	wsIx.unsub(jdata, ch)
}

func (wsIx *WSIndexClient) ReqBasis(contractCode string, period string, callbackFun OnReqBasisResponse, from int64, to int64,
	basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	reqData := wsbase.WSReqData{Req: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(reqData)

	wsIx.req(jdata, ch, callbackFun, reflect.TypeOf(index.ReqBasisResponse{}))
}

func (wsIx *WSIndexClient) UnreqBasis(contractCode string, period string, from int64, to int64, basisPriceType string, id string) {
	if basisPriceType == "" {
		basisPriceType = "open"
	}
	if id == "" {
		id = linearswap.DEFAULT_ID
	}
	ch := fmt.Sprintf("market.%s.basis.%s.%s", contractCode, period, basisPriceType)
	unreqData := wsbase.WSUnreqData{Unreq: ch, Id: id, From: from, To: to}
	jdata, _ := json.Marshal(unreqData)

	wsIx.unreq(jdata, ch)
}
