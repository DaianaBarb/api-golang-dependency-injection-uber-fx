package repository

import (
	"context"
	"database/sql"
	"fmt"
	model "golang-uber-fx/core/domain"
	"golang-uber-fx/util/errors"
	"log"
	"strings"
)

//criar logs

type IUserRepository interface {
	SaveUser(cliente *model.User) error
	FindUser(cpf string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) SaveUser(user *model.User) error {
	query := `CREATE TABLE IF NOT EXISTS user( user_username VARCHAR(50), user_password VARCHAR(100))`
	_, err := r.db.Exec(query)
	if err != nil {
		log.Printf("Error %s when creating user table", err)

	}
	query2 := `CREATE TABLE IF NOT EXISTS client_cli(client_name VARCHAR(50) PRIMARY KEY NOT NULL UNIQUE, client_tel VARCHAR(50), client_cpf VARCHAR(50) , client_createdAt DATE, client_active boolean DEFAULT false )`
	_, err = r.db.Exec(query2)
	if err != nil {
		log.Printf("Error %s when creating client_cli table", err)
	}

	query3 := `CREATE TABLE IF NOT EXISTS adress_cli(
  zipCode varchar(20) NOT NULL,
  publicPlace varchar(45) NOT NULL,
  neighborhood varchar(45) NOT NULL,
  location varchar(45) NOT NULL,
  uf varchar(45) NOT NULL,
  state varchar(45) NOT NULL,
  region varchar(45) NOT NULL,
  cpf_cli char(20) NOT NULL,
  adress_id int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (adress_id),
  KEY fk_cpf_cli_idx (cpf_cli),
  CONSTRAINT fk_cpf_cli FOREIGN KEY (cpf_cli) REFERENCES client_cli (client_cpf)
) ENGINE=InnoDB AUTO_INCREMENT=1010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;`
	_, err = r.db.Exec(query3)
	if err != nil {
		log.Printf("Error %s when creating adress_cli table", err)
	}
	query4 := `SET autocommit = ON;`
	_, err = r.db.Exec(query4)
	if err != nil {
		log.Printf("Error %s when cSET autocommit = ON;", err)
	}
	_, err = r.db.Exec("INSERT INTO user (user_username, user_password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.AlreadyExistsf("Name '%s' create error: user already exists", user.Username)
		} else {
			return fmt.Errorf("user save error, name:  %s: error: %v", user.Username, err)
		}
	}
	return nil
}

func (r *UserRepository) FindUser(name string) (*model.User, error) {

	cli := model.User{}
	err := r.db.QueryRowContext(context.TODO(), "SELECT user_username, user_password FROM user WHERE user_username = ?", name).Scan(&cli.Username, &cli.Password)
	//if err := row.Scan(&cli);
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundf("name %s: not found user", name)
		}
		return nil, fmt.Errorf("name  %s: error: %v", name, err)
	}
	return &cli, nil
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}

}
