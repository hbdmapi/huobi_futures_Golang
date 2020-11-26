package ws

import (
	"container/list"
	"encoding/json"
	"fmt"
	"huobi_futures_Golang/sdk/linearswap"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/sdk/reqbuilder"
	"huobi_futures_Golang/sdk/wsbase"
	"reflect"
	"time"

	"github.com/gorilla/websocket"
)

type MethonInfo struct {
	fun   interface{}
	param reflect.Type
}

func (wsOp *MethonInfo) Init(fun interface{}, param reflect.Type) *MethonInfo {
	wsOp.fun = fun
	wsOp.param = param
	return wsOp
}

type WebSocketOp struct {
	host string
	path string

	conn      *websocket.Conn
	accessKey string
	secretKey string

	autoConnect bool
	allSubStrs  list.List

	onSubCallbackFuns map[string]*MethonInfo
	onReqCallbackFuns map[string]*MethonInfo

	authOk bool
}

func (wsOp *WebSocketOp) open(path string, host string, accessKey string, secretKey string, autoConnect bool) bool {
	if host == "" {
		wsOp.host = linearswap.LINEAR_SWAP_DEFAULT_HOST
	}

	wsOp.host = host
	wsOp.path = path

	wsOp.accessKey = accessKey
	wsOp.secretKey = secretKey

	wsOp.autoConnect = autoConnect

	ret := wsOp.connServer()
	if ret {
		wsOp.allSubStrs = list.List{}
		wsOp.onSubCallbackFuns = make(map[string]*MethonInfo)
		wsOp.onReqCallbackFuns = make(map[string]*MethonInfo)
	}
	return ret
}

func (wsOp *WebSocketOp) close() {
	wsOp.conn.Close()
	wsOp.conn = nil
}

func (wsOp *WebSocketOp) connServer() bool {
	url := fmt.Sprintf("wss://%s%s", wsOp.host, wsOp.path)
	var err error
	wsOp.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Error("WebSocket connected error: %s", err)
		return false
	}
	log.Info("WebSocket connected")
	wsOp.conn.SetCloseHandler(wsOp.onClose)

	go wsOp.readLoop(wsOp.conn)

	if wsOp.accessKey == "" && wsOp.secretKey == "" {
		wsOp.authOk = true
		return true
	}
	wsOp.authOk = false
	return wsOp.sendAuth(wsOp.conn, wsOp.host, wsOp.path, wsOp.accessKey, wsOp.secretKey)
}

func (wsOp *WebSocketOp) onClose(code int, text string) error {
	log.Info("WebSocket close.")

	wsOp.conn = nil
	if wsOp.autoConnect {
		if !wsOp.connServer() {
			return fmt.Errorf("reconnect server error")
		}
	}

	return fmt.Errorf("")
}

func (wsOp *WebSocketOp) sendAuth(conn *websocket.Conn, host string, path string, accessKey string, secretKey string) bool {
	if conn == nil {
		log.Error("websocket conn is null")
		return false
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	req := new(reqbuilder.GetRequest).Init()
	req.AddParam("AccessKeyId", accessKey)
	req.AddParam("SignatureMethod", "HmacSHA256")
	req.AddParam("SignatureVersion", "2")
	req.AddParam("Timestamp", timestamp)

	sign := new(reqbuilder.Signer).Init(secretKey)
	signature := sign.Sign("GET", host, path, req.BuildParams())

	auth := wsbase.WSAuthData{
		Op:               "auth",
		AtType:           "api",
		AccessKeyId:      accessKey,
		SignatureMethod:  "HmacSHA256",
		SignatureVersion: "2",
		Timestamp:        timestamp,
		Signature:        signature}

	jdata, error := json.Marshal(&auth)
	if error != nil {
		log.Error("Auth to json error.")
		return false
	}
	conn.WriteMessage(websocket.TextMessage, jdata)
	return true
}

func (wsOp *WebSocketOp) readLoop(conn *websocket.Conn) {
	for conn != nil {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			log.Error("Read error: %s", err)
			continue
		}
		var message string
		if msgType == websocket.BinaryMessage {
			message, err = wsbase.GZipDecompress(buf)
			if err != nil {
				log.Error("UnGZip data error: %s", err)
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = string(buf)
		}

		var jdata map[string]interface{}
		err = json.Unmarshal([]byte(message), &jdata)
		if err != nil {
			log.Error("msg to map[string]json.RawMessage error: %s", err)
			continue
		}

		if ts, found := jdata["ping"]; found { // market heartbeat
			ts = int64(ts.(float64))
			//log.Info("WebSocket received data: \"ping\":%d", ts)

			pongData := fmt.Sprintf("{\"pong\":%d }", ts)
			wsOp.conn.WriteMessage(websocket.TextMessage, []byte(pongData))

			//log.Info("WebSocket replied data: %s", pongData)
		} else if op, found := jdata["op"]; found {
			switch op {
			case "ping": // order heartbeat
				ts := jdata["ts"]
				//log.Info("WebSocket received data, { \"op\":\"%s\", \"ts\": \"%s\" }", op, ts)

				pongData := fmt.Sprintf("{ \"op\":\"pong\", \"ts\": \"%s\" }", ts)
				wsOp.conn.WriteMessage(websocket.TextMessage, []byte(pongData))

				//log.Info("WebSocket replied data, %s", pongData)

			case "close":
				log.Error("Some error occurres when authentication in server side.")

			case "error":
				log.Error("Illegal op or internal error, but websoket is still connected.")

			case "auth":
				code := int64(jdata["err-code"].(float64))
				if code == 0 {
					log.Info("Authentication success.")
					wsOp.authOk = true
					for e := wsOp.allSubStrs.Front(); e != nil; e = e.Next() {
						wsOp.conn.WriteMessage(websocket.TextMessage, []byte(e.Value.(string)))
					}
				} else {
					msg := jdata["err-msg"].(string)
					log.Error("Authentication failure: %d/%s", code, msg)
					wsOp.close()
				}
			case "notify":
				topic := jdata["topic"].(string)
				wsOp.handleSubCallbackFun(topic, message, jdata)
			case "sub":
				topic := jdata["topic"]
				log.Info("sub: \"%s\"", topic)
			case "unsub":
				topic := jdata["topic"].(string)
				log.Info("unsub: \"%s\"", topic)

				if _, found := wsOp.onSubCallbackFuns[topic]; found {
					delete(wsOp.onSubCallbackFuns, topic)
				}
			default:
				log.Info("WebSocket received unknow data: %s", jdata)
			}
		} else if topic, found := jdata["subbed"]; found { // sub success reply
			log.Info("\"subbed\": \"%s\"", topic)
		} else if topic, found := jdata["unsubbed"]; found { // unsub success reply
			log.Info("\"unsubbed\": \"%s\"", topic)

			if _, found := wsOp.onSubCallbackFuns[topic.(string)]; found {
				delete(wsOp.onSubCallbackFuns, topic.(string))
			}
		} else if topic, found := jdata["ch"]; found { // market sub reply data

			wsOp.handleSubCallbackFun(topic.(string), message, jdata)
		} else if topic, found := jdata["rep"]; found { // market request reply data
			wsOp.handleReqCallbackFun(topic.(string), message, jdata)
		} else if code, found := jdata["err-code"]; found { // market request reply data
			code = code
			msg := jdata["err-msg"]
			log.Error("%d:%s", code, msg)
		} else {
			log.Info("WebSocket received unknow data: %s", jdata)
		}
	}
}

func (wsOp *WebSocketOp) handleSubCallbackFun(ch string, data string, jdata map[string]interface{}) {
	var mi *MethonInfo = nil
	if _, found := wsOp.onSubCallbackFuns[ch]; found {
		mi = wsOp.onSubCallbackFuns[ch]

	} else if ch == "accounts" || ch == "positions" {
		//contract_code := jdata["data"][0]["contract_code"], &contract_code)
		contract_code := jdata["data"]

		full_ch := fmt.Sprintf("%s.%s", ch, contract_code)
		if _, found := wsOp.onSubCallbackFuns[full_ch]; found {
			mi = wsOp.onSubCallbackFuns[full_ch]
		} else if _, found := wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]; found {
			mi = wsOp.onSubCallbackFuns[fmt.Sprintf("%s.*", ch)]
		}
	} else if ch[:7] == "orders." {
		if _, found := wsOp.onSubCallbackFuns["orders.*"]; found {
			mi = wsOp.onSubCallbackFuns["orders.*"]
		}
	} else if ch[:12] == "matchOrders." {
		if _, found := wsOp.onSubCallbackFuns["matchOrders.*"]; found {
			mi = wsOp.onSubCallbackFuns["matchOrders.*"]
		}
	} else if ch[:14] == "trigger_order." {
		if _, found := wsOp.onSubCallbackFuns["trigger_order.*"]; found {
			mi = wsOp.onSubCallbackFuns["trigger_order.*"]
		}
	} else if ch[len(ch)-19:] == ".liquidation_orders" {
		if _, found := wsOp.onSubCallbackFuns["public.*.liquidation_orders"]; found {

			mi = wsOp.onSubCallbackFuns["public.*.liquidation_orders"]
		}
	}

	if mi == nil {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	wsOp.runFunction(mi, data)
}

func (wsOp *WebSocketOp) handleReqCallbackFun(ch string, data string, jdata map[string]interface{}) {
	var mi *MethonInfo = nil
	var found bool = false

	if mi, found = wsOp.onReqCallbackFuns[ch]; !found {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	if mi == nil {
		log.Error("no callback function to handle: %s", jdata)
		return
	}

	wsOp.runFunction(mi, data)
}

func (wsOp *WebSocketOp) runFunction(mi *MethonInfo, data string) {
	param := reflect.New(mi.param).Interface()
	json.Unmarshal([]byte(data), &param)

	rargs := make([]reflect.Value, 1)
	rargs[0] = reflect.ValueOf(param)

	fun := reflect.ValueOf(mi.fun)
	fun.Call(rargs)
}

func (wsOp *WebSocketOp) sub(subStr []byte, ch string, fun interface{}, param reflect.Type) bool {

	for !wsOp.authOk {
		time.Sleep(10)
	}

	var mi *MethonInfo = nil
	var found bool
	if mi, found = wsOp.onSubCallbackFuns[ch]; found {
		mi = new(MethonInfo).Init(fun, param)
		wsOp.onSubCallbackFuns[ch] = mi
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, subStr)
	log.Info("websocket has send data: %s", subStr)

	wsOp.allSubStrs.PushBack(string(subStr))

	mi = new(MethonInfo).Init(fun, param)
	wsOp.onSubCallbackFuns[ch] = mi

	return true
}

func (wsOp *WebSocketOp) unsub(unsubStr []byte, ch string) bool {

	for !wsOp.authOk {
		time.Sleep(10)
	}

	if _, found := wsOp.onSubCallbackFuns[ch]; !found {
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, unsubStr)
	log.Info("websocket has send data: %s", unsubStr)

	var next *list.Element
	for e := wsOp.allSubStrs.Front(); e != nil; e = next {
		next = e.Next()
		val := e.Value.(string)
		if val == string(unsubStr) {
			wsOp.allSubStrs.Remove(e)
		}
	}

	return true
}

func (wsOp *WebSocketOp) req(subStr []byte, ch string, fun interface{}, param reflect.Type) bool {

	for !wsOp.authOk {
		time.Sleep(10)
	}

	var mi *MethonInfo = nil
	var found bool
	if mi, found = wsOp.onReqCallbackFuns[ch]; found {
		mi = new(MethonInfo).Init(fun, param)
		wsOp.onReqCallbackFuns[ch] = mi
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, subStr)
	log.Info("websocket has send data: %s", subStr)

	mi = new(MethonInfo).Init(fun, param)
	wsOp.onReqCallbackFuns[ch] = mi

	return true
}

func (wsOp *WebSocketOp) unreq(unsubStr []byte, ch string) bool {

	for !wsOp.authOk {
		time.Sleep(10)
	}

	if _, found := wsOp.onReqCallbackFuns[ch]; !found {
		return true
	}
	wsOp.conn.WriteMessage(websocket.TextMessage, unsubStr)
	log.Info("websocket has send data: %s", unsubStr)

	return true
}
