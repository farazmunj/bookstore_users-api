package services

import "github.com/farazmunj/bookstore_users-api/domain/users"

//CreateUser save user in DB
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
