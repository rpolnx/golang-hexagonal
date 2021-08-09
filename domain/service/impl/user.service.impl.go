package service

import (
	"time"

	v "gopkg.in/dealancer/validate.v2"

	errs "github.com/pkg/errors"
	ce "rpolnx.com.br/mongo-hex/application/error"
	u "rpolnx.com.br/mongo-hex/domain/model/user"
	"rpolnx.com.br/mongo-hex/domain/ports/out"
	"rpolnx.com.br/mongo-hex/domain/service"
)

type userService struct {
	repository out.UserRepository
}

func (u *userService) Get(name string) (*u.User, error) {
	return u.repository.Get(name)
}

func (u *userService) Post(user *u.User) error {
	if err := v.Validate(user); err != nil {
		return errs.Wrap(ce.ErrInvalid, "service.User.Post")
	}
	user.CreatedAt = time.Now()
	return u.repository.Post(user)
}

func (u *userService) Delete(name string) error {
	return u.repository.Delete(name)
}

func NewUserService(repo out.UserRepository) service.UserService {
	return &userService{
		repo,
	}
}
