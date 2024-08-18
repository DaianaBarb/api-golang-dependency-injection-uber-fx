package main

import (
	"fmt"
	"os"

	"golang-uber-fx/fx/server"

	"github.com/joho/godotenv"
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
	server.Start2()

}
