package service

import (
	repository "golang-uber-fx/adapter/mysql/repository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
)

type IUserService interface {
	SaveUser(user *model.User) error
	FindUser(name string) (*model.User, error)
}

type UserService struct {
	repository repository.IUserRepository
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


func (s *UserService) SaveUser(user *model.User) error {
	// criar DTO e transdormar aqui
	err := s.repository.SaveUser(user)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return err

	}
	s.log.LogLevelInfo("------------save sucess-------------")
	return nil
}

func NewUserService(repo repository.IUserRepository, l log.ILogLevel) IUserService {
	return &UserService{
		repository: repo,
		log:        l,
	}

}
