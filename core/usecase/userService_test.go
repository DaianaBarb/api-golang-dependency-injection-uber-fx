package service

import (
	repository "golang-uber-fx/adapter/mysql/repository"
	model "golang-uber-fx/core/domain"
	log "golang-uber-fx/util/log"
	"reflect"
	"testing"
)

func TestUserService_SaveUser(t *testing.T) {
	type fields struct {
		repository repository.IUserRepository
		log        log.ILogLevel
	}
	type args struct {
		user *model.User
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
			s := &UserService{
				repository: tt.fields.repository,
				log:        tt.fields.log,
			}
			if err := s.SaveUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.SaveUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_FindUser(t *testing.T) {
	type fields struct {
		repository repository.IUserRepository
		log        log.ILogLevel
	}
	type args struct {
		cpf string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &UserService{
				repository: tt.fields.repository,
				log:        tt.fields.log,
			}
			got, err := s.FindUser(tt.args.cpf)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.FindUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.FindUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
