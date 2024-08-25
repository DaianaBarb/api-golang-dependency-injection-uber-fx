package dto

import (
	model "golang-uber-fx/core/domain"
	"time"
)

type ClientDtoResponse struct {
	Name      string    `json:"name"`
	Tel       string    `json:"tel"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
}

type ClientDtoRequest struct {
	Name string `json:"name"`
	Tel  string `json:"tel"`
	Cpf  string `json:"cpf"`
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
	}

}
