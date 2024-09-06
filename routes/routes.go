package routes

import (
	"fmt"
	hand "golang-uber-fx/adapter/http"
	"net/http"

	"github.com/gorilla/mux"
)

type IRoutes interface {
	RegisterRoutes()
}

func (s *Routes) RegisterRoutes() {

	c := mux.NewRouter()

	c.HandleFunc("/client/{cpf}", s.hanCli.Find).Methods("GET")
	c.HandleFunc("/client", s.hanCli.Save).Methods("POST")
	c.HandleFunc("/client", s.hanCli.FindAllByParam).Methods("GET").Queries()
	c.HandleFunc("/client/{cpf}", s.hanCli.Del).Methods("DELETE")
	c.HandleFunc("/user", s.hanUse.Save).Methods("POST")
	c.HandleFunc("/user/{name}", s.hanUse.Find).Methods("GET")
	c.HandleFunc("/user/passport", s.hanUse.CreateToken).Methods("POST")

	fmt.Println(" online na porta 8080")
	http.ListenAndServe(":8080", c)
}

type Routes struct {
	hanCli hand.IClientServer
	hanUse hand.IUserServer
}

func NewRoutes(hanCli hand.IClientServer, hanUse hand.IUserServer) IRoutes {

	return &Routes{
		hanCli: hanCli,
		hanUse: hanUse,
	}

}
