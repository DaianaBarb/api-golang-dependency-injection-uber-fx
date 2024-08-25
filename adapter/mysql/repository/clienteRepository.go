package repository

import (
	"database/sql"
	"fmt"
	model "golang-uber-fx/core/domain"
	"golang-uber-fx/util/errors"
	"strings"
)

//criar logs

type IClientRepository interface {
	SaveClient(client *model.Client) error
	DeleteClient(cpf string) error
	FindClient(cpf string) (*model.Client, error)
}

type ClientRepository struct {
	db *sql.DB
}

// DeleteCliente implements Irepository.
func (r *ClientRepository) DeleteClient(cpf string) error {

	_, err := r.db.Exec("DELETE FROM client_cli WHERE client_cpf=?", cpf)
	if err != nil {
		return fmt.Errorf("client delete error, cpf:  %s: error: %v", cpf, err)

	}
	return nil
}

// SaveCliente implements IClientRepository.
func (r *ClientRepository) SaveClient(client *model.Client) error {
	_, err := r.db.Exec("INSERT INTO client_cli (client_name, client_tel, client_cpf, client_createdAt, client_active ) VALUES (?, ?, ?, ?, ?)", client.Name, client.Tel, client.Cpf, client.CreatedAt, client.Active)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.AlreadyExistsf("Name '%s' create error: client already exists", client.Name)
		} else {
			return fmt.Errorf("cliente save error, name:  %s: error: %v", client.Name, err)
		}
	}
	return nil
}

func (r *ClientRepository) FindClient(cpf string) (*model.Client, error) {

	cli := &model.Client{}
	row := r.db.QueryRow("SELECT client_name, client_tel, client_cpf, client_createdAt, client_active FROM client_cli WHERE client_cpf = ?", cpf)
	if err := row.Scan(&cli); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundf("cpf %s: not found client", cpf)
		}
		return nil, fmt.Errorf("cpf  %s: error: %v", cpf, err)
	}
	return cli, nil
}

func NewClientRepository(db *sql.DB) IClientRepository {
	return &ClientRepository{
		db: db,
	}

}
