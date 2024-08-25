package dto

import (
	model "golang-uber-fx/core/domain"
)

type UserDTO struct {
	Username string `json:"userName"`
	Token    string `json:"token"`
}

func ToUserDto(u *model.User) *UserDTO {

	return &UserDTO{
		Username: u.Username,
	}

}
