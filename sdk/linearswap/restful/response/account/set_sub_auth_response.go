package account

type SetSubAuthResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		Errors []struct {
			SubUid string `json:"sub_uid,omitempty"`

			ErrorCode int `json:"err_code,omitempty"`

			ErrorMessage string `json:"err_msg,omitempty"`
		} `json:"errors"`
		Successes string `json:"successes"`
	} `json:"data,omitempty"`
	Ts int64 `json:"ts"`
}
