package account

type GetFinancialRecordExactResponse struct {
	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data struct {
		FinancialRecord []struct {
			Id int64 `json:"id"`

			Ts int64 `json:"ts"`

			Asset string `json:"asset"`

			ContractCode string `json:"contract_code"`

			MarginAccount string `json:"margin_account"`

			FaceMarginAccount string `json:"face_margin_account"`

			FcType int `json:"type,omitempty"`

			Amount float32 `json:"amount"`

			// endregion
		} `json:"financial_record"`

		RemainSize int `json:"remain_size"`

		NextId int64 `json:"next_id,omitempty"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
