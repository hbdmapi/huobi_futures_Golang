package wsbase

type WSAuthData struct {
	Op string `json:"op"`

	AtType string `json:"type"`

	Cid string `json:"cid,omitempty"`

	AccessKeyId string `json:"AccessKeyId"`

	SignatureMethod string `json:"SignatureMethod"`

	SignatureVersion string `json:"SignatureVersion"`

	Timestamp string `json:"Timestamp"`

	Signature string `json:"Signature"`

	Ticket string `json:"ticket,omitempty"`
}

func (wsAuthData *WSAuthData) Init() *WSAuthData {
	wsAuthData.Op = "auth"
	wsAuthData.AtType = "api"
	wsAuthData.SignatureMethod = "HmacSHA256"
	wsAuthData.SignatureVersion = "2"

	return wsAuthData
}
