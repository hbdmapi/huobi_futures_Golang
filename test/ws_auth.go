package main

import (
	"bytes"
	"compress/gzip"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

func readLoop(conn *websocket.Conn) {
	for conn != nil {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Read error: %s", err)
			continue
		}
		var message string
		if msgType == websocket.BinaryMessage {
			message, err = gzipDecompress(buf)
			if err != nil {
				fmt.Printf("UnGZip data error: %s", err)
				continue
			}
		} else if msgType == websocket.TextMessage {
			message = string(buf)
		}
		fmt.Println(message)
	}
}

func sign(sk string, method string, host string, path string, parameters string) string {

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	sb.WriteString(parameters)

	fmt.Println("sign parameters:", parameters)

	payload := sb.String()
	hash := hmac.New(sha256.New, []byte(sk))
	hash.Write([]byte(payload))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}

func sendAuth(conn *websocket.Conn, host string, path string, accessKey string, secretKey string) bool {
	timestamp := time.Now().UTC().Format("2006-01-02T15:04:05")

	urls := url.Values{}
	urls.Add("AccessKeyId", accessKey)
	urls.Add("SignatureMethod", "HmacSHA256")
	urls.Add("SignatureVersion", "2")
	urls.Add("Timestamp", timestamp)

	signature := sign(secretKey, "GET", host, path, urls.Encode())
	fmt.Println("signature:", signature)

	auth := fmt.Sprintf("{\"op\":\"auth\", \"type\":\"api\",\"AccessKeyId\":\"%s\", \"SignatureMethod\":\"HmacSHA256\", \"SignatureVersion\":\"2\", \"Timestamp\":\"%s\", \"Signature\":\"%s\"}", accessKey, timestamp, signature)
	fmt.Println("send auth:", auth)
	jdata := []byte(auth)
	conn.WriteMessage(websocket.TextMessage, jdata)
	return true
}

func gzipDecompress(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	reader, gzipErr := gzip.NewReader(buf)
	if gzipErr != nil {
		return "", gzipErr
	}
	defer reader.Close()

	result, readErr := ioutil.ReadAll(reader)
	if readErr != nil {
		return "", readErr
	}

	return string(result), nil
}

func main() {
	host := "api.hbdm.com"
	path := "/swap-notification"
	ak := "x-x-x-x"
	sk := "x-x-x-x"

	url := fmt.Sprintf("wss://%s%s", host, path)
	fmt.Println("url:", url)

	//head := http.Header{"Host": {host + ":8080"}}
	//conn, _, err := websocket.DefaultDialer.Dial(url, head)
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		fmt.Println("WebSocket connected error: %s", err)
	} else {
		//fmt.Println("WebSocket connected")
		sendAuth(conn, host, path, ak, sk)
		readLoop(conn)
	}
}
