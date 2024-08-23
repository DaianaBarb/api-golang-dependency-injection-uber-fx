package service

import (
	repository "golang-uber-fx/adapter/mysql/userRepository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
)

type IUserservice interface {
	SaveUser(user *model.User) error
	FindUser(name string) (*model.User, error)
}

type UserService struct {
	repository repository.Irepository
	log        log.ILogLevel
}

// FindCliente implements Iservice.
func (s *UserService) FindUser(cpf string) (*model.User, error) {
	// criar DTO e transdormar aqui

	user, err := s.repository.FindUser(cpf)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return nil, err

	}
	s.log.LogLevelInfo("------------save sucess-------------")

	return user, err
}
