package services

import (
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userid int64) (*users.User, *errors.RestErr) {
	u := users.User{Id: userid}
	if err := u.Get(); err != nil {
		return nil, err
	}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}


func UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	curUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	curUser.FirstName = user.FirstName
	curUser.LastName = user.LastName
	curUser.Email = user.Email

	if err := curUser.Update(); err != nil {
		return nil, err
	}
	
	return curUser, nil
}
