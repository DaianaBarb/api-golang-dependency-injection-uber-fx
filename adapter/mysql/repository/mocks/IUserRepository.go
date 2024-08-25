// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import domain "golang-uber-fx/core/domain"
import mock "github.com/stretchr/testify/mock"

// IUserRepository is an autogenerated mock type for the IUserRepository type
type IUserRepository struct {
	mock.Mock
}

// FindUser provides a mock function with given fields: cpf
func (_m *IUserRepository) FindUser(cpf string) (*domain.User, error) {
	ret := _m.Called(cpf)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(cpf)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(cpf)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: cliente
func (_m *IUserRepository) SaveUser(cliente *domain.User) error {
	ret := _m.Called(cliente)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(cliente)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}