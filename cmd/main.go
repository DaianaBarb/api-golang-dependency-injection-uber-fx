package main

import (
	"client/internal/api/rest"
	"client/internal/repository/mysql"
	"client/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

const (
	ENV_FILE = "../config/.dev.env"
)

func init() {

	if _, err := os.Stat(ENV_FILE); os.IsNotExist(err) {
		fmt.Println("error in env file")
	}

	err := godotenv.Load(ENV_FILE)
	if err != nil {
		fmt.Println("error in env file")

	}

}

func main() {
	fmt.Println("iniciando....")

	app := fx.New(
		mysql.Module,
		service.Module,
		rest.Module,
		fx.Invoke(func(handler rest.IClientHandler) error {

			r := mux.NewRouter()
			r.HandleFunc("/api/client", handler.CreatedClientHandler).Methods("POST")
			go http.ListenAndServe(":8080", r)
			fmt.Println("on na porta 8080")

			return nil

		}),
	)
	app.Run()
}
