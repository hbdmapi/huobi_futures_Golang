package account

type GetAccountTransHisResponse struct {
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

			Amount float32 `json:"amount"`

			// region only for account
			FaceMarginAccount string `json:"face_margin_account,omitempty"`

			FcType int `json:"type,omitempty"`
			// endregion

			// region only for sub account
			FromMarginAccount string `json:"from_margin_account,omitempty"`

			ToMarginAccount string `json:"to_margin_account,omitempty"`

			SubUid string `json:"sub_uid,omitempty"`

			SubAccountName string `json:"sub_account_name,omitempty"`

			TransferType int `json:"transfer_type,omitempty"`
			// endregion
		} `json:"financial_record"`

		TotalPage int `json:"total_page"`

		CurrentPage int `json:"current_page"`

		TotalSize int `json:"total_size"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
