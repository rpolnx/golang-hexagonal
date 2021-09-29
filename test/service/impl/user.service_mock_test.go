package service

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	ce "rpolnx.com.br/golang-hex/application/error"
	u "rpolnx.com.br/golang-hex/domain/model/user"
	"rpolnx.com.br/golang-hex/domain/ports/out"
	"time"
)

type mongoRepositorySuccess struct {
	mongoRepository out.MongoUserRepository
}

func (*mongoRepositorySuccess) FindAll() ([]u.User, error) {
	return []u.User{
		{
			ID:        primitive.NewObjectID(),
			Name:      "Bob",
			Age:       20,
			Birthday:  "20/20/2000",
			CreatedAt: time.Now(),
		},
		{
			ID:        primitive.NewObjectID(),
			Name:      "Alice",
			Age:       25,
			Birthday:  "20/20/1996",
			CreatedAt: time.Now(),
		},
	}, nil

}
func (*mongoRepositorySuccess) Find(name string) (*u.User, error) {
	return &u.User{
		ID:        primitive.NewObjectID(),
		Name:      "Alice",
		Age:       25,
		Birthday:  "20/20/1996",
		CreatedAt: time.Now(),
	}, nil

}
func (*mongoRepositorySuccess) Post(person *u.User) error {
	return nil
}
func (*mongoRepositorySuccess) Delete(name string) error {
	return nil
}

type cacheRepositorySuccess struct {
	cacheRepository out.RedisUserRepository
}

func (*cacheRepositorySuccess) Get(name string) (*u.User, error) {
	return &u.User{
		ID:        primitive.NewObjectID(),
		Name:      name,
		Age:       25,
		Birthday:  "20/20/1996",
		CreatedAt: time.Now(),
	}, nil
}

func (*cacheRepositorySuccess) Set(_ *u.User) error {
	return nil
}

type cacheRepositorySetFail struct {
	cacheRepository out.RedisUserRepository
}

func (*cacheRepositorySetFail) Get(name string) (*u.User, error) {
	return nil, ce.ErrCacheMiss
}

func (*cacheRepositorySetFail) Set(_ *u.User) error {
	return errors.New("set error")
}
