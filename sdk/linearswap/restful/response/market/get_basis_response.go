package market

type GetBasisResponse struct {
	Ch string `json:"ch"`

	Status string `json:"status"`

	ErrorCode int `json:"err_code,omitempty"`

	ErrorMessage string `json:"err_msg,omitempty"`

	Data []struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"data,omitempty"`

	Ts int64 `json:"ts"`
}
