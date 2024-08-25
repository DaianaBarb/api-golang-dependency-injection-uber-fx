package http

import (
	"encoding/json"
	"fmt"
	"golang-uber-fx/core/dto"
	service "golang-uber-fx/core/usecase"
	"golang-uber-fx/util"
	"golang-uber-fx/util/errors"
	"net/http"

	"github.com/gorilla/mux"
)

type IClientServer interface {
	Save(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	Del(w http.ResponseWriter, r *http.Request)
}

type ClientServer struct {
	serv service.IClientService
}

//criar logs

func NewServer(serv service.IClientService) IClientServer {

	return &ClientServer{
		serv: serv,
		//router: &rout,
	}

}

func (c *ClientServer) Save(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := json.NewEncoder(w).Encode("Missing authorization header")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode("Unauthorized")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	cliRequest := new(dto.ClientDtoRequest)
	err = json.NewDecoder(r.Body).Decode(&cliRequest)
	if err != nil {

		errors.UnprocessableEntityf("unprocessable entity error: %v", err)
		return
	}
	util.ValidateStruct(cliRequest)

	c.serv.SaveClient(dto.ToClientModel(cliRequest))
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
func (c *ClientServer) Find(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := json.NewEncoder(w).Encode("Missing authorization header")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode("Unauthorized")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}

	//vars := r.URL.Query()  este modo pega as queries da url, do outro modo pega o {cpf} da url definida
	vars := mux.Vars(r)
	cpf, ok := vars["cpf"]
	if !ok {
		errors.NewBadRequest(nil, "error to parser query params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cli, err := c.serv.FindClient(cpf)
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if cli != nil {
		err = json.NewEncoder(w).Encode(cli)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	}
	w.WriteHeader(http.StatusNotFound)

}
func (c *ClientServer) Del(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := json.NewEncoder(w).Encode("Missing authorization header")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode("Unauthorized")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)

		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return
	}
	vars := mux.Vars(r)
	cpf, ok := vars["cpf"]
	if !ok {
		errors.NewBadRequest(nil, "error to parser query params")
		return
	}
	c.serv.DeleteClient(cpf)
	w.WriteHeader(http.StatusOK)

}
