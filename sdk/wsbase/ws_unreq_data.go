package wsbase

type WSUnreqData struct {
	Unreq string `json:"unreq"`

	Id string `json:"id"`

	From int64 `json:"from"`

	To int64 `json:"to"`
}
