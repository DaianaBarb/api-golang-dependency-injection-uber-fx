package service

import (
	"client/internal/domain/entities"
	"client/internal/domain/models"
	"client/internal/repository/mysql"
	"time"
)

type IclientService interface {
	CreatedClient(cli *models.Client) error
}

type clientService struct {
	repository mysql.IClientRepository
}

func NewClientService(repo mysql.IClientRepository) IclientService {
	return &clientService{
		repository: repo,
	}
}

func (c *clientService) CreatedClient(i *models.Client) error {

	err := c.repository.CreatedClient(&entities.Client{Id: i.Id,
		Name:      i.Name,
		Active:    i.Active,
		Telefone:  i.Telefone,
		CreatedAt: time.Now()})
	if err != nil {
		return err
	}

	return nil
}
