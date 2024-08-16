package domain

type Cliente struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Cpf  string `json:"cpf"`
}

type ClienteFindQuery struct {
	Cpf string `query:"cpf"`
}
