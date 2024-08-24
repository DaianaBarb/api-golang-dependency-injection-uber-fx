package repository

import (
	"database/sql"
	"fmt"
	model "golang-uber-fx/core/domain"
	"golang-uber-fx/util/errors"
	"strings"
)

type IUserRepository interface {
	SaveUser(cliente *model.User) error
	FindUser(cpf string) (*model.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) SaveUser(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO user (user_username, user_password,) VALUES (?, ?)", user.Username, user.Password)
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

	cli := &model.User{}
	row := r.db.QueryRow("SELECT user_username, user_password FROM user WHERE user_username = ?", name)
	if err := row.Scan(&cli); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundf("name %s: not found user", name)
		}
		return nil, fmt.Errorf("name  %s: error: %v", name, err)
	}
	return cli, nil
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}

}
