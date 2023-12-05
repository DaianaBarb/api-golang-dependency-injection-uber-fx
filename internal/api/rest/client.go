package rest

import (
	"client/internal/domain/dto"
	"client/internal/domain/models"
	"client/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	Service service.IclientService
}

type IClientHandler interface {
	CreatedClientHandler(w http.ResponseWriter, r *http.Request)
}

func NewCLientHandler(serv service.IclientService) IClientHandler {
	return &ClientHandler{
		Service: serv,
	}
}

func (c *ClientHandler) CreatedClientHandler(w http.ResponseWriter, r *http.Request) {

	var clientRequest dto.ClientRequest
	//ctx := context.WithValue(context.Background(),  uuid.New().String(),  uuid.New().String())
	err := json.NewDecoder(r.Body).Decode(&clientRequest)
	if err != nil {

		fmt.Println("unprocessable entity")
		return
	}

	error := validateStruct(&clientRequest)
	if error != nil {
		fmt.Println("bad request")

		return
	}

	err = c.Service.CreatedClient(&models.Client{
		Name:     clientRequest.ClientName,
		Active:   clientRequest.ClientActive,
		Telefone: clientRequest.ClientTel,
	})
	if err != nil {

		w.WriteHeader(500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return

}

func validateStruct(v interface{}) error {
	var validate *validator.Validate
	validate = validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
