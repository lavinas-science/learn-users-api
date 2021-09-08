package services

import "github.com/lavinas-science/learn-users-api/domain/users"


func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}

