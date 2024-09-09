package http

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	model "golang-uber-fx/core/domain"
	dto "golang-uber-fx/core/dto"
	service "golang-uber-fx/core/usecase"
	"golang-uber-fx/util"
	"golang-uber-fx/util/errors"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/mux"
)

type IUserServer interface {
	Save(w http.ResponseWriter, r *http.Request)
	Find(w http.ResponseWriter, r *http.Request)
	CreateToken(w http.ResponseWriter, r *http.Request)
}

type UserServer struct {
	serv service.IUserService
}

// criar logs
func (u *UserServer) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		errors.NewBadRequest(nil, "error to parser query params")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := u.serv.FindUser(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func createHashSha256(password string) string {

	h := sha256.New()
	h.Write([]byte(password))
	p := h.Sum(nil)
	return hex.EncodeToString(p)
}
func (u *UserServer) Save(w http.ResponseWriter, r *http.Request) {

	userRequest := new(model.User)
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {

		errors.UnprocessableEntityf("unprocessable entity error: %v", err)
		return
	}
	util.ValidateStruct(userRequest)
	userRequest.Password = createHashSha256(userRequest.Password)
	err = u.serv.SaveUser(userRequest)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	token, err := CreateToken(userRequest.Username)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = json.NewEncoder(w).Encode(&dto.UserDTO{
		Username: userRequest.Username,
		Token:    token,
	})
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return

}

func NewUserServer(serv service.IUserService) IUserServer {

	return &UserServer{
		serv: serv,
	}

}

func (c *UserServer) CreateToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var u model.User
	json.NewDecoder(r.Body).Decode(&u)
	user, err := c.serv.FindUser(u.Username)
	if err != nil {

		if errors.IsNotFound(err) {
			w.WriteHeader(http.StatusNotFound)

		}
		w.WriteHeader(http.StatusInternalServerError)

	}
	if user != nil {
		if createHashSha256(u.Password) == user.Password {
			token, err := CreateToken(user.Username)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

			}
			json.NewEncoder(w).Encode(token)

		} else {
			json.NewEncoder(w).Encode("senha inválida")

			w.WriteHeader(http.StatusUnprocessableEntity)
		}

	}

}

// criar um endpoint NewUser pra criar o usuario novo e ja gerar o token na hora
// utilizar este endpoint pra criar token quando expirar depois de 24 horas
// assim que receber o login validar no banco o usuario se ele existe e se a senha e valida, se tudo tiver valido gerar o token se n for valida retornar erro
// fazer uma busca por nome se encontrar validar a senha, se for valida gerar o token
// salvar a senha criptografada no banco e descriptografar quando retornar

func CreateToken(username string) (string, error) {
	// futuramente salvar token em uma coluna temporária do usuario com data e hora e verificar se o token possui menos de 24 horas, caso ele tenha nao gerar outro token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username":   username,
			"exp":        time.Now().Add(time.Hour * 24).Unix(),
			"authorized": true,
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
