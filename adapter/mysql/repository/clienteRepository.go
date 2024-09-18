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
	FindAllClientByParam(name, tel, cpf, active, createdAt string, limit int, page int) ([]model.Client, *model.PaginationData, error)
}

const (
	SELECT_CLIENT       = "SELECT  client_name, client_tel, client_cpf, client_createdAt, client_active  FROM client_cli "
	SELECT_CLIENT_COUNT = "SELECT count(*) FROM client_cli "
)

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
	// _, err := r.db.Exec("INSERT INTO client_cli (client_name, client_tel, client_cpf, client_createdAt, client_active ) VALUES (?, ?, ?, ?, ?)", client.Name, client.Tel, client.Cpf, client.CreatedAt, client.Active)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "already exists") {
	// 		return errors.AlreadyExistsf("Name '%s' create error: client already exists", client.Name)
	// 	} else {
	// 		return fmt.Errorf("cliente save error, name:  %s: error: %v", client.Name, err)
	// 	}
	// }

	_, err := r.db.Exec("START TRANSACTION; INSERT INTO client_cli (client_name, client_tel, client_cpf, client_createdAt, client_active ) VALUES (?, ?, ?, ?, ?); INSERT INTO adress_cli(zipCode,publicPlace,neighborhood,location,uf,state,region,cpf_cli) VALUES (?, ?, ?, ?, ?,?,?,?); COMMIT;", client.Name, client.Tel, client.Cpf, client.CreatedAt, client.Active, client.Address.ZipCode, client.Address.PublicPlace, client.Address.Neighborhood, client.Address.Location, client.Address.Uf, client.Address.State, client.Address.Region, client.Cpf)
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

	cli := &model.Client{
		Address: &model.Address{},
	}
	row := r.db.QueryRow("SELECT C.client_name, C.client_tel, C.client_cpf, C.client_createdAt,C.client_active, A.zipCode,A.publicPlace,A.neighborhood,A.location,A.uf,A.state,A.region from client_cli as C  join adress_cli as A ON C.client_cpf = A.cpf_cli WHERE client_cpf = ?", cpf)
	if err := row.Scan(&cli.Name, &cli.Tel, &cli.Cpf, &cli.CreatedAt, &cli.Active, &cli.Address.ZipCode, &cli.Address.PublicPlace, &cli.Address.Neighborhood, &cli.Address.Location, &cli.Address.Uf, &cli.Address.State, &cli.Address.Region); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.NotFoundf("cpf %s: not found client", cpf)
		}
		return nil, fmt.Errorf("cpf  %s: error: %v", cpf, err)
	}
	return cli, nil
}

func (r *ClientRepository) FindAllClientByParam(name, tel, cpf, active, createdAt string, limit int, page int) ([]model.Client, *model.PaginationData, error) {
	itemList := []model.Client{}
	pagination := &model.PaginationData{}
	offset := limit * (page - 1)

	var count int = 0
	var args []any
	sqlWhere := ""
	conectName := " WHERE "

	if name != "" {
		sqlWhere = sqlWhere + conectName + "client_name=?"
		conectName = " AND "

		args = append(args, name)

	}
	if tel != "" {
		sqlWhere = sqlWhere + conectName + "client_tel=?"
		conectName = " AND "

		args = append(args, tel)

	}
	if cpf != "" {
		sqlWhere = sqlWhere + conectName + "client_cpf=?"
		conectName = " AND "

		args = append(args, cpf)

	}
	if active != "" {
		sqlWhere = sqlWhere + conectName + "client_active=?"
		conectName = " AND "
		if active == "true" {
			args = append(args, 1)
		} else {
			args = append(args, 0)
		}
	}
	if createdAt != "" {
		sqlWhere = sqlWhere + conectName + "client_createdAt=?"
		conectName = " AND "

		args = append(args, createdAt)

	}
	rows, err := r.db.Query(SELECT_CLIENT_COUNT+sqlWhere, args...)
	if err != nil {
		return nil, nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			defer rows.Close()
			return nil, nil, err
		}
	}
	rows.Close()
	orderBy := " ORDER BY client_name LIMIT ? OFFSET ?"
	args = append(args, limit)
	args = append(args, offset)
	rows, err = r.db.Query(SELECT_CLIENT+sqlWhere+orderBy, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		tag, err := scanRowClient(rows)
		if err != nil {
			return nil, nil, err
		}
		itemList = append(itemList, *tag)
	}

	n := count % limit
	nOfPages := count / limit
	if n > 0 {
		nOfPages = nOfPages + 1
	}

	pagination.Page = page
	pagination.TotalPage = nOfPages
	pagination.Limit = limit
	pagination.Total = count
	pagination.Count = len(itemList)

	return itemList, pagination, nil

}
func scanRowClient(rows *sql.Rows) (*model.Client, error) {
	item := new(model.Client)
	err := rows.Scan(&item.Name, &item.Tel, &item.Cpf,
		&item.CreatedAt, &item.Active)

	return item, err
}

func NewClientRepository(db *sql.DB) IClientRepository {
	return &ClientRepository{
		db: db,
	}

}
