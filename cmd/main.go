package main

import (
	"fmt"
	"golang-uber-fx/fx/server"
	"os"

	"github.com/joho/godotenv"
)

const (
	ENV_FILE = "../config/.dev.env"
)

func init() {

	// prometheus.MustRegister(pro.DurationSeconds)
	// prometheus.MustRegister(pro.ResponseStatus)

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

// 	h := sha256.New()
// 	h.Write([]byte("12345"))
// 	p := h.Sum(nil)
// 	fmt.Println(hex.EncodeToString(p))

// 	if "5994471abb01112afcc18159f6cc74b4f511b99806da59b3caf5a9c173cacfc5" == hex.EncodeToString(p) {

// 		fmt.Println("ok")
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
// 		jwt.MapClaims{
// 			"username":   "daiana",
// 			"exp":        time.Now().Add(time.Hour * 24).Unix(),
// 			"authorized": true,
// 		})

// 	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
// 	if err != nil {

// 	}

// 	fmt.Println("token")

// 	fmt.Println(tokenString)
// 	err = verifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQ1NDEyNzUsInVzZXJuYW1lIjoiZGFpYW5hIn0.TMGCwD9IeiyWBRiguqyYI6teizf8sC0xT8Cb3PVBuzs")
// 	if err != nil {

// 	}
// }
// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return fmt.Errorf("invalid token")
// 	}

// 	return nil
// }
