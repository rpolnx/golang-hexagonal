package out

import u "rpolnx.com.br/golang-hex/domain/model/user"

type MongoUserRepository interface {
	FindAll() ([]u.User, error)
	Find(name string) (*u.User, error)
	Post(person *u.User) error
	Delete(name string) error
}

type RedisUserRepository interface {
	Get(name string) (*u.User, error)

	Set(user *u.User) error
}
