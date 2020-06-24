package services

import (
	"github.com/farazmunj/bookstore_users-api/domain/users"
	"github.com/farazmunj/bookstore_users-api/utils/errors"
)

//CreateUser save user in DB
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
