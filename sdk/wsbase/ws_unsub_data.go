package wsbase

type WSUnsubData struct {
	Unsub string `json:"unsub"`

	Id string `json:"id"`

	DataType string `json:"data_type,omitempty"`
}
