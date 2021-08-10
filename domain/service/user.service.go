package service

import (
	u "rpolnx.com.br/golang-hex/domain/model/user"
)

type UserService interface {
	GetAll() ([]u.User, error)
	Get(name string) (*u.User, error)
	Post(*u.User) error
	Delete(name string) error
}
