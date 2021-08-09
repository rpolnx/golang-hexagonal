package service

import (
	u "rpolnx.com.br/golang-hex/domain/model/user"
)

type UserService interface {
	Get(name string) (*u.User, error)
	Post(*u.User) error
	Delete(name string) error
}
