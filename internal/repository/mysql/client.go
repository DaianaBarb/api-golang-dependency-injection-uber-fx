package mysql

import (
	"client/internal/domain/entities"
	"database/sql"
	"fmt"
	"strings"

	"github.com/americanas-go/errors"
)

type ClientRepository struct {
	db *sql.DB
}

type IClientRepository interface {
	CreatedClient(cli *entities.Client) error
}

func NewClient(db *sql.DB) IClientRepository {
	return &ClientRepository{
		db: db,
	}
}

func (c *ClientRepository) CreatedClient(cli *entities.Client) error {

	sqlResult, err := c.db.Exec("INSERT INTO client (client_name, active, client_tel, created_at) VALUES (?, ?, ?, ?)", cli.Name, cli.Active, cli.Telefone, cli.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			return errors.AlreadyExistsf("tagName '%s' create error: client already exists", cli.Name)
		} else {
			return fmt.Errorf("tagName '%s' create error: %v", cli.Name, err)
		}
	}
	newId, err := sqlResult.LastInsertId()
	cli.Id = int(newId)

	return nil

}
