package service

import (
	repository "golang-uber-fx/adapter/mysql/repository"
	model "golang-uber-fx/core/domain"
	"golang-uber-fx/core/dto"
	log "golang-uber-fx/util/log"
)

type IClientService interface {
	SaveClient(cliente *model.Client) error
	DeleteClient(cpf string) error
	FindClient(cpf string) (*dto.ClientDtoResponse, error)
	FindAllClientByParam(name, tel, cpf, active, createdAt string, limit int, page int) ([]dto.ClientDtoResponse, *model.PaginationData, error)
}

type Service struct {
	repository repository.IClientRepository
	log        log.ILogLevel
}

// FindCliente implements Iservice.
func (s *Service) FindClient(cpf string) (*dto.ClientDtoResponse, error) {
	// criar DTO e transdormar aqui

	cli, err := s.repository.FindClient(cpf)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return nil, err

	}
	s.log.LogLevelInfo("------------save sucess-------------")

	return dto.ToClientDTOResponse(cli), err
}

// DeleteCliente implements Iservice.
func (s *Service) DeleteClient(cpf string) error {
	// criar DTO e transdormar aqui
	err := s.repository.DeleteClient(cpf)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return err

	}
	s.log.LogLevelError("delete error")
	return nil
}

func (s *Service) SaveClient(client *model.Client) error {

	err := s.repository.SaveClient(client)
	if err != nil {
		s.log.LogLevelError("------------save error-------------")
		return err

	}
	s.log.LogLevelInfo("------------save sucess-------------")
	return nil
}
func (s *Service) FindAllClientByParam(name, tel, cpf, active, createdAt string, limit int, page int) ([]dto.ClientDtoResponse, *model.PaginationData, error) {
	modelList, pag, err := s.repository.FindAllClientByParam(name, tel, cpf, active, createdAt, limit, page)
	if err != nil {
		return nil, nil, err
	}
	newListDto := dto.ToClientListResponse(modelList)

	return newListDto, pag, nil
}

func NewService(repo repository.IClientRepository, l log.ILogLevel) IClientService {
	return &Service{
		repository: repo,
		log:        l,
	}

}
