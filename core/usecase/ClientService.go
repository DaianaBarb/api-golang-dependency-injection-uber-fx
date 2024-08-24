package service

import (
	repository "golang-uber-fx/adapter/mysql/repository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
)

type Iservice interface {
	SaveCliente(cliente *model.Cliente) error
	DeleteCliente(cpf string) error
	FindCliente(cpf string) (*model.Cliente, error)
}

type Service struct {
	repository repository.IClientRepository
	log        log.ILogLevel
}

// FindCliente implements Iservice.
func (s *Service) FindCliente(cpf string) (*model.Cliente, error) {
	// criar DTO e transdormar aqui

	cliente, err := s.repository.FindCliente(cpf)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return nil, err

	}
	s.log.LogLevelInfo("------------save sucess-------------")

	return cliente, err
}

// DeleteCliente implements Iservice.
func (s *Service) DeleteCliente(cpf string) error {
	// criar DTO e transdormar aqui
	err := s.repository.DeleteCliente(cpf)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return err

	}
	s.log.LogLevelError("delete error")
	return nil
}

// SaveCliente implements Iservice.
func (s *Service) SaveCliente(cliente *model.Cliente) error {
	// criar DTO e transdormar aqui
	err := s.repository.SaveCliente(cliente)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return err

	}
	s.log.LogLevelInfo("------------save sucess-------------")
	return nil
}

func NewService(repo repository.IClientRepository, l log.ILogLevel) Iservice {
	return &Service{
		repository: repo,
		log:        l,
	}

}
