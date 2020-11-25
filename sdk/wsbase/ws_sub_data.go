package wsbase

type WSSubData struct {
	Sub string `json:"sub"`

	Id string `json:"id"`

	DataType string `json:"data_type,omitempty"`
}
