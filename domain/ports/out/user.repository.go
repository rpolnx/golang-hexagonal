package out

import u "rpolnx.com.br/mongo-hex/domain/model/user"

type UserRepository interface {
	Get(name string) (*u.User, error)
	Post(person *u.User) error
	Delete(name string) error
}
