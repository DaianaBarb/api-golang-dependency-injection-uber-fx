package service

import (
	repository "golang-uber-fx/adapter/mysql/clienteRepository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
)

type Iservice interface {
	SaveCliente(cliente *model.Cliente)
	DeleteCliente(cpf string)
	FindCliente(cpf string)
}

type Service struct {
	repository repository.Irepository
	log        log.ILogLevel
}

// FindCliente implements Iservice.
func (s *Service) FindCliente(cpf string) {
	s.repository.FindCliente(cpf)
}

// DeleteCliente implements Iservice.
func (s *Service) DeleteCliente(cpf string) {
	s.repository.DeleteCliente(cpf)
	s.log.LogLevelError("delete error")

}

// SaveCliente implements Iservice.
func (s *Service) SaveCliente(cliente *model.Cliente) {

	s.repository.SaveCliente(cliente)
	s.log.LogLevelInfo("------------save sucess-------------")
}

func NewService(repo repository.Irepository, l log.ILogLevel) Iservice {
	return &Service{
		repository: repo,
		log:        l,
	}

}
