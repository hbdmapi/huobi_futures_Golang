package wsbase

type WSOpData struct {
	Op string `json:"op"`

	Cid string `json:"cid,omitempty"`

	Topic string `json:"topic"`
}
