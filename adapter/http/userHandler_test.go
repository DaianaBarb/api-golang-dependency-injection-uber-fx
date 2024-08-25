package http

import (
	service "golang-uber-fx/core/usecase"
	"net/http"
	"testing"
)

func TestUserServer_Find(t *testing.T) {
	type fields struct {
		serv service.IUserService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserServer{
				serv: tt.fields.serv,
			}
			u.Find(tt.args.w, tt.args.r)
		})
	}
}

func TestUserServer_Save(t *testing.T) {
	type fields struct {
		serv service.IUserService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserServer{
				serv: tt.fields.serv,
			}
			u.Save(tt.args.w, tt.args.r)
		})
	}
}

func TestUserServer_CreateToken(t *testing.T) {
	type fields struct {
		serv service.IUserService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &UserServer{
				serv: tt.fields.serv,
			}
			c.CreateToken(tt.args.w, tt.args.r)
		})
	}
}
