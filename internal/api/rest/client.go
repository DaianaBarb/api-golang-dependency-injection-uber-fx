package rest

import (
	"client/internal/domain/dto"
	"client/internal/domain/models"
	"client/internal/service"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/americanas-go/errors"
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

type ResponseDefault struct {
	Message string `json:"message"`
}

func (c *ClientHandler) CreatedClientHandler(w http.ResponseWriter, r *http.Request) {

	var clientRequest dto.ClientRequest
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
		if errors.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)
			encodeErr := json.NewEncoder(w).Encode(ResponseDefault{Message: err.Error()})
			if encodeErr != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func validateStruct(v interface{}) error {

	validate := validator.New()

	errs := validate.Struct(v)
	if errs != nil {
		return errs
	}
	return nil
}
