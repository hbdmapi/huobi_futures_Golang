package index

type SubBasiesResponse struct {
	Ch string `json:"rep"`

	Ts int64 `json:"ts"`

	Tick struct {
		Id int64 `json:"id"`

		ContractPrice string `json:"contract_price"`

		IndexPrice string `json:"index_price"`

		Basis string `json:"basis"`

		BasisRate string `json:"basis_rate"`
	} `json:"tick"`
}
