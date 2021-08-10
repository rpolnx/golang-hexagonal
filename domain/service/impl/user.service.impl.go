package service

import (
	"time"

	v "gopkg.in/dealancer/validate.v2"

	errs "github.com/pkg/errors"
	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/ports/out"
	"rpolnx.com.br/golang-hex/domain/service"
)

type userService struct {
	mongoRepository out.MongoUserRepository
	cacheRepository out.RedisUserRepository
}

func (u *userService) GetAll() ([]u.User, error) {
	return u.mongoRepository.FindAll()
}

func (u *userService) Get(name string) (*u.User, error) {
	user, err := u.cacheRepository.Get(name)

	if err == nil {
		return user, nil
	}

	if err != ce.ErrCacheMiss {
		return nil, err
	}

	created, err := u.mongoRepository.Find(name)

	if err != nil {
		return nil, err
	}

	err = u.cacheRepository.Set(created)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *userService) Post(user *u.User) error {
	if err := v.Validate(user); err != nil {
		return errs.Wrap(ce.ErrInvalid, "service.User.Post")
	}
	user.CreatedAt = time.Now()
	return u.mongoRepository.Post(user)
}

func (u *userService) Delete(name string) error {
	return u.mongoRepository.Delete(name)
}

func NewUserService(mongoRepo out.MongoUserRepository, redisRepo out.RedisUserRepository) service.UserService {
	return &userService{
		mongoRepository: mongoRepo,
		cacheRepository: redisRepo,
	}
}
