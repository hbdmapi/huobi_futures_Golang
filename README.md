# Huobi Futures Golang SDK (but only supports linearswap now)

This is Huobi Golang SDK, you can import it to your package.

The SDK API supports both RESTful and websocket to get/sub the market, account and order infomation.

## Table of Contents

- [Quick Start](#Quick-Start)

- [Usage](#Usage)

  - [Configuration](#Configuration)
  - [Folder structure](#Folder-Structure)
  - [Client](#Client)
  - [Response](#Response)
  
- [Request examples](#Request-examples)
  - [Market data](#Market-data)

- [Subscription examples](#Subscription-examples)
  - [Subscribe trade update](#Subscribe-trade-update)

  

## Quick Start

The SDK is run and test in go1.15.4, you can import the source code from remote or local what shuold be dowload first.

The package **sdk** is core code source as SDK API.
The package **test** is a unit test what you can find usage of each API interface in.

If you want to create your own application, you can follow below steps:

* Create a client (xxxClient/WSxxxClient) instance.
* Call the method of client instance.

```go
// RESTful api
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")

data := make(chan account.GetAccountAssetsResponse)
go acClient.GetAccountAssetsAsync(data, "BTC-USDT", 0)
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// success
}

// websocker api
wsmkClient := new(ws.WSMarketClient).Init("")
wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
	// show response data with *m
}, "")
time.Sleep(time.Duration(10) * time.Second)
wsmkClient.UnsubKLine("BTC-USDT", "15min", "")
time.Sleep(time.Duration(10) * time.Second)
```

## Usage

### Configuration

The client Init function must set AccessKey/SecretKey two value when you access private data. And it not need to set AccessKey/SecretKey value when you access public data such as market data.

You can create config.json in your package for config AccessKey/SecretKey and other input data.

```json
{
  "Host": "api.hbdm.com",
  "AccessKey": "x-x-x-x",
  "SecretKey": "x-x-x-x",
  "AccountId": 10000000,
  "SubUid": 10000000
}
```
Then create config.go file to read the config.json file

```go
import (
	"encoding/json"
	"os"
)

type Config struct {
	Host      string
	AccessKey string
	SecretKey string
	AccountId int64
	SubUid    int64
}

var config *Config

func init() {

	filePtr, err := os.Open("config.json")
	if err != nil {
		return
	}
	defer filePtr.Close()

	config = new(Config)
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(config)
}
 ```

And use it as follow:
```go
accessKey := config.AccessKey
secretKey := config.SecretKey
```

### Folder Structure

This is the folder and namespace structure of SDK source code and the description

- **sdk**: The SDK API package
  - **linearswap**: the linear swap api src inclue RESTful and websocket
  - **requestbuilder**: Responsible to build the request with the signature
  - **log**: The internal logger interface and implementations
  - **wsbase**: The websocket data model
- **test**: The unit test package
  - **xxx_test.go**: The golang unit test file

You can find all demo in xxx_test.go to get/sub linear swap private/public data

### Client

In this SDK, the client is the object to access the Huobi API. All the client are listed in below table. Each client is very small and simple.

| Access Type | Client | Privacy | Data Category  |
| ----------- | -------| ------- | ------------ |
| RESTful     | AccountClient | Private | account info |
|             | MarketClient | Public | market info |
|             | OrderClient | Private | about order |
|             | TransferClient | Private | transfer assets |
|             | TriggerOrderClient | Private | about trigger order |
| Websocket   | WSIndexClinet | Public | index info |
|             | WSMarketClinet | Public | market info |
|             | WSNotifyClinet | Public/Private | market info/ account info |
|             |                |         |              |

#### Public vs. Private

There are two types of privacy that is correspondent with privacy of API:

**Public client**: It invokes public API to get public data (Common data and Market data), therefore you can create a new instance without applying an API Key.

```go
mkClient := restful.MarketClient{}
mkClient.Init("")
```

**Private client**: It invokes private API to access private data, you need to follow the API document to apply an API Key first, and pass the API Key to the Init.

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")
```

The API key is used for authentication. If the authentication cannot pass, the invoking of private interface will fail.

#### Rest vs. WebSocket

**Rest Client**: It invokes Rest API and get once-off response.

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")

data := make(chan account.GetAccountAssetsResponse)
go acClient.GetAccountAssetsAsync(data, "BTC-USDT", 0)
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// success
}
```

**WebSocket Client**: It establishes WebSocket connection with server and data will be pushed from server actively. There are two types of method for WebSocket client:

- Request method: The method name starts with "Req", it will receive the once-off data after sending the request.
- Subscription: The method name starts with "Sub", it will receive update after sending the subscription.

```go
wsmkClient := new(ws.WSMarketClient).Init("")

// Reqxxx
wsmkClient.ReqKLine("BTC-USDT", "1min", func(m *market.ReqKLineResponse) {
	// show response data with *m
}, 1604395758, 1604396758, "")

// Subxxx
wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
	// show response data with *m
}, "")
time.Sleep(time.Duration(10) * time.Second)
wsmkClient.UnsubKLine("BTC-USDT", "15min", "")
time.Sleep(time.Duration(10) * time.Second)
```

#### Custom host

Each client Init support an optional host parameter, by default it is "api.btcgateway.pro". If you need to use different host, you can specify the custom host. 

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "Host")

wsmkClient := new(ws.WSMarketClient).Init("Host")
```

### Response

All response data are defined as follow:
- **huobi_futures_Golang.sdk.linearswap.restful.response**: all RESTful response data
- **huobi_futures_Golang.sdk.linearswap.ws.response**: all websockt response data

## Request Examples
### Market data
```go
mkClient := restful.MarketClient{}
mkClient.Init("")

data := make(chan market.GetContractInfoResponse)

go mkClient.GetContractInfoAsync(data, "")
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// show data with x
}
```

## Subscription Examples
### Subscribe trade update
```go
wsnfClient := new(ws.WSNotifyClient).Init("AccessKey", "SecretKey", "")

wsnfClient.SubOrders("*", func(m *notify.SubOrdersResponse) {
	// show dat with *m
}, "")
time.Sleep(time.Duration(50) * time.Second)
wsnfClient.UnsubOrders("*", "")
time.Sleep(time.Duration(50) * time.Second)
```