package domain

import "time"

type Client struct {
	Name      string    `json:"name"`
	Tel       string    `json:"tel"`
	Cpf       string    `json:"cpf"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

type ClienteFindQuery struct {
	Cpf string `query:"cpf"`
}
