package services

import (
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.Status = users.StatusActive
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

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	curUser, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if !isPartial || user.FirstName != "" {
		curUser.FirstName = user.FirstName
	}
	if !isPartial || user.LastName != "" {
		curUser.LastName = user.LastName
	}
	if !isPartial || user.Email != "" {
		curUser.Email = user.Email
	}
	if err := curUser.Validate(); err != nil {
		return nil, err
	}
	if err := curUser.Update(); err != nil {
		return nil, err
	}
	return curUser, nil
}

func DeleteUser(userId int64) *errors.RestErr {
	usr := &users.User{Id: userId}
	if err := usr.Delete(); err != nil {
		return err
	}
	return nil
}

func Search(status string) ([]users.User, *errors.RestErr) {
	d := &users.User{}
	return d.FindByStatus(status)
}
