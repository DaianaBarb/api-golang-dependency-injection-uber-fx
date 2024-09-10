package http

import (
	"encoding/json"
	"fmt"
	"golang-uber-fx/core/domain"
	"golang-uber-fx/core/dto"
	service "golang-uber-fx/core/usecase"
	"golang-uber-fx/util"
	"golang-uber-fx/util/errors"
	pro "golang-uber-fx/util/observability"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type IClientServer interface {
	Save(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	Del(w http.ResponseWriter, r *http.Request)
	FindAllByParam(w http.ResponseWriter, r *http.Request)
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

	defer func() {

		requestDuration := pro.DurationSeconds
		totalRequests := pro.TotalRequests
		start := time.Now()
		duration := time.Since(start)

		requestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())
		totalRequests.WithLabelValues(r.URL.Path).Inc()
	}()
	responseStatus := pro.ResponseStatus
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		err := json.NewEncoder(w).Encode("Missing authorization header")
		if err != nil {
			responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusInternalServerError)).Inc()
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode("Unauthorized")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusInternalServerError)).Inc()

			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusUnauthorized)).Inc()
		return
	}

	tokenString = tokenString[len("Bearer "):]

	err := VerifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusUnauthorized)).Inc()
		fmt.Fprint(w, "Invalid token")
		return
	}

	cliRequest := new(dto.ClientDtoRequest)
	err = json.NewDecoder(r.Body).Decode(&cliRequest)
	if err != nil {
		responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusUnprocessableEntity)).Inc()

		errors.UnprocessableEntityf("unprocessable entity error: %v", err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	util.ValidateStruct(cliRequest)

	err = c.serv.SaveClient(dto.ToClientModel(cliRequest))
	if err != nil {
		responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusInternalServerError)).Inc()
		w.WriteHeader(http.StatusInternalServerError)

	}
	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusInternalServerError)).Inc()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	responseStatus.WithLabelValues(r.URL.Path, strconv.Itoa(http.StatusOK)).Inc()


}
func (c *ClientServer) Find(w http.ResponseWriter, r *http.Request) {
	defer func() {

		responseStatus := pro.ResponseStatus
		requestDuration := pro.DurationSeconds
		totalRequests := pro.TotalRequests
		start := time.Now()
		duration := time.Since(start)
		responseStatus.WithLabelValues(r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())
		totalRequests.WithLabelValues(r.URL.Path).Inc()
	}()

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
	defer func() {

		responseStatus := pro.ResponseStatus
		requestDuration := pro.DurationSeconds
		totalRequests := pro.TotalRequests
		start := time.Now()
		duration := time.Since(start)
		responseStatus.WithLabelValues(r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())
		totalRequests.WithLabelValues(r.URL.Path).Inc()
	}()

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

func (c *ClientServer) FindAllByParam(w http.ResponseWriter, r *http.Request) {
	defer func() {

		responseStatus := pro.ResponseStatus
		requestDuration := pro.DurationSeconds
		totalRequests := pro.TotalRequests
		start := time.Now()
		duration := time.Since(start)
		responseStatus.WithLabelValues(r.URL.Path).Inc()
		requestDuration.WithLabelValues(r.URL.Path).Observe(duration.Seconds())
		totalRequests.WithLabelValues(r.URL.Path).Inc()
	}()
	w.Header().Set("Content-Type", "application/json")

	cpf := r.URL.Query().Get("cpf")

	name := r.URL.Query().Get("name")

	active := r.URL.Query().Get("active")

	createdAt := r.URL.Query().Get("createdAt")

	tel := r.URL.Query().Get("tel")

	lim := r.URL.Query().Get("limit")

	limit := 0

	if lim == "" {
		limit = 30
	} else {

		n, err := strconv.Atoi(lim)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		limit = n
		if limit > 100 {
			limit = 100
		}

	}

	pg := r.URL.Query().Get("page")

	if pg == "" {
		pg = "1"
	}

	n, err := strconv.Atoi(pg)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	page := n

	items, pagination, err := c.serv.FindAllClientByParam(name, tel, cpf, active, createdAt, limit, page)

	if err != nil {

		if errors.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)

			return
		}

		return

	}

	response := dto.ResponsePagination{
		Items: items,
		Result: domain.PaginationData{
			TotalPage: pagination.TotalPage,
			Count:     pagination.Count,
			Page:      pagination.Page,
			Limit:     pagination.Limit,
			Total:     pagination.Total,
		},
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
