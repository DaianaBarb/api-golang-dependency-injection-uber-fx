package service

import (
	repository "golang-uber-fx/adapter/mysql/clienteRepository"
	model "golang-uber-fx/core/domain"
)

type Iservice interface {
	SaveCliente(cliente *model.Cliente)
	DeleteCliente(cpf string)
	FindCliente(cpf string)
}

type Service struct {
	repository repository.Irepository
}

// FindCliente implements Iservice.
func (s *Service) FindCliente(cpf string) {
	s.repository.FindCliente(cpf)
}

// DeleteCliente implements Iservice.
func (s *Service) DeleteCliente(cpf string) {
	s.repository.DeleteCliente(cpf)
}

// SaveCliente implements Iservice.
func (s *Service) SaveCliente(cliente *model.Cliente) {

	s.repository.SaveCliente(cliente)
}

func NewService(repo repository.Irepository) Iservice {
	return &Service{
		repository: repo,
	}

}
