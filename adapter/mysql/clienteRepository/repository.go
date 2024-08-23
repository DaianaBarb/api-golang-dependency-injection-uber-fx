package repository

import (
	"database/sql"
	"fmt"
	model "golang-uber-fx/core/domain"
	"golang-uber-fx/util/errors"
	"strings"
)

type Irepository interface {
	SaveCliente(cliente *model.Cliente) error
	DeleteCliente(cpf string) error
	FindCliente(cpf string) (*model.Cliente, error)
}

type Repository struct {
	db *sql.DB
}

// DeleteCliente implements Irepository.
func (r *Repository) DeleteCliente(cpf string) error {

	_, err := r.db.Exec("DELETE FROM cliente WHERE cliente_cpf=?", cpf)
	if err != nil {
		return fmt.Errorf("cliente delete error, cpf:  %s: error: %v", cpf, err)

	}
	return nil
}

// SaveCliente implements Irepository.
func (r *Repository) SaveCliente(cliente *model.Cliente) error {
	_, err := r.db.Exec("INSERT INTO cliente (client_name, client_tel, client_cpf, ) VALUES (?, ?, ?)", cliente.Name, cliente.Tel, cliente.Cpf)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.AlreadyExistsf("Name '%s' create error: client already exists", cliente.Name)
		} else {
			return fmt.Errorf("cliente save error, name:  %s: error: %v", cliente.Name, err)
		}
	}
	return nil
}

func (r *Repository) FindCliente(cpf string) (*model.Cliente, error) {

	cli := &model.Cliente{}
	row := r.db.QueryRow("SELECT client_name, client_tel, client_cpf FROM cliente WHERE client_cpf = ?", cpf)
	if err := row.Scan(&cli); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundf("cpf %s: not found tag", cpf)
		}
		return nil, fmt.Errorf("cpf  %s: error: %v", cpf, err)
	}
	return cli, nil
}

func NewRepository(db *sql.DB) Irepository {
	return &Repository{
		db: db,
	}

}
