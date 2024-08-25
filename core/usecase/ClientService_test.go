package service

import (
	repository "golang-uber-fx/adapter/mysql/repository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
	"testing"
)

func TestService_SaveClient(t *testing.T) {
	type fields struct {
		repository repository.IClientRepository
		log        log.ILogLevel
	}
	type args struct {
		client *model.Client
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
				log:        tt.fields.log,
			}
			if err := s.SaveClient(tt.args.client); (err != nil) != tt.wantErr {
				t.Errorf("Service.SaveClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteClient(t *testing.T) {
	type fields struct {
		repository repository.IClientRepository
		log        log.ILogLevel
	}
	type args struct {
		cpf string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				repository: tt.fields.repository,
				log:        tt.fields.log,
			}
			if err := s.DeleteClient(tt.args.cpf); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
