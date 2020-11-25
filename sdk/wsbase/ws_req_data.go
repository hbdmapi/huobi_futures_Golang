package wsbase

type WSReqData struct {
	Req string `json:"req"`

	Id string `json:"id"`

	From int64 `json:"from,omitempty"`

	To int64 `json:"to,omitempty"`
}
