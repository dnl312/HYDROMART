package repo

import "auth/model"

type AuthInterface interface {
	LoginUser(user model.User) (string, error)
	RegisterUser(user model.RegisterUser) error
}