package domain

import "time"

type Client struct {
	Name      string    `json:"name"`
	Tel       string    `json:"tel"`
	Cpf       string    `json:"cpf"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	Address   *Address   `json:"address"`
}

type Address struct {
	ZipCode      string `json:"zipCode"`
	PublicPlace  string `json:"publicPlace"`
	Neighborhood string `json:"neighborhood"`
	Location     string `json:"location"`
	Uf           string `json:"uf"`
	State        string `json:"state"`
	Region       string `json:"region"`
}

type ClienteFindQuery struct {
	Cpf string `query:"cpf"`
}
type PaginationData struct {
	TotalPage int `json:"totalPage"`
	Count     int `json:"count"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	Total     int `json:"total"`
}
