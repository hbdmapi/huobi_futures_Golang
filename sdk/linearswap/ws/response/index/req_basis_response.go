package index

type ReqBasisResponse struct {
	Rep string `json:"rep"`

	Status string `json:"status"`

	Id string `json:"id"`

	Wsid int64 `json:"wsid"`

	Ts int64 `json:"ts"`

	Data []struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"data"`
}
