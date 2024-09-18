package dto

import (
	"golang-uber-fx/core/domain"
	model "golang-uber-fx/core/domain"
	"time"
)

type ClientDtoResponse struct {
	Name      string         `json:"name"`
	Tel       string         `json:"tel"`
	Active    bool           `json:"active"`
	CreatedAt time.Time      `json:"createdAt"`
	Adress    domain.Address `json:"adress"`
}

type ClientDtoRequest struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Cpf  string `json:"cpf"`
	Cep  string `json:"cep"`
}

type ResponsePagination struct {
	Items  []ClientDtoResponse  `json:"items"`
	Result model.PaginationData `json:"_result"`
}

func ToClientModel(cli *ClientDtoRequest) *model.Client {

	return &model.Client{
		Name:      cli.Name,
		Tel:       cli.Tel,
		Cpf:       cli.Cpf,
		Active:    true,
		CreatedAt: time.Now(),
	}

}

func ToClientDTOResponse(cli *model.Client) *ClientDtoResponse {

	return &ClientDtoResponse{
		Name:      cli.Name,
		Tel:       cli.Tel,
		Active:    cli.Active,
		CreatedAt: cli.CreatedAt,
		Adress: model.Address{ZipCode: cli.Address.ZipCode,
			PublicPlace:  cli.Address.PublicPlace,
			Neighborhood: cli.Address.Neighborhood,
			Location:     cli.Address.Location,
			Uf:           cli.Address.Uf,
			State:        cli.Address.State,
			Region:       cli.Address.Region},
	}

}

func ToClientListResponse(cli []model.Client) []ClientDtoResponse {

	list := []ClientDtoResponse{}

	for _, c := range cli {

		list = append(list, ClientDtoResponse{
			Name:      c.Name,
			Tel:       c.Tel,
			Active:    c.Active,
			CreatedAt: c.CreatedAt})

	}

	return list

}
