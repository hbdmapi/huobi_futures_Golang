package account

type GetAccountTransHisResponse struct {
	status string

	errorCode string `json:"err_code,omitempty"`

	errorMessage string `json:"err_msg,omitempty"`

	data []struct {
		financialRecord []struct {
			id int64

			ts int64

			asset string

			contractCode string `json:"contract_code"`

			marginAccount string `json:"margin_account"`

			amount float32

			// region only for account
			faceMarginAccount string `json:"face_margin_account,omitempty"`

			fcType int `json:"type,omitempty"`
			// endregion

			// region only for sub account
			fromMarginAccount string `json:"from_margin_account,omitempty"`

			toMarginAccount string `json:"to_margin_account,omitempty"`

			subUid string `json:"sub_uid,omitempty"`

			subAccountName string `json:"sub_account_name,omitempty"`

			transferType int `json:"transfer_type,omitempty"`
			// endregion
		} `json:"financial_record"`

		totalPage int64 `json:"total_page"`

		currentPage int64 `json:"current_page"`

		totalSize int64 `json:"total_size"`
	} `json:"omitempty"`

	ts int64
}
