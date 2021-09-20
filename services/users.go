package services

import (
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-utils-go/crypto"
	"github.com/lavinas-science/learn-utils-go/dates"
	"github.com/lavinas-science/learn-utils-go/rest_errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	GetUser(int64) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	SearchUser(string) (users.Users, *rest_errors.RestErr)
	LoginUser(users.Login) (*users.User, *rest_errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *rest_errors.RestErr) {
	user.Status = users.StatusActive
	user.DateCreated = dates.GetNowDb()
	user.Password = crypto.GetMd5(user.Password)

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userid int64) (*users.User, *rest_errors.RestErr) {
	u := users.User{Id: userid}
	if err := u.Get(); err != nil {
		return nil, err
	}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *rest_errors.RestErr) {
	curUser, err := UserService.GetUser(user.Id)
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

func (s *userService) DeleteUser(userId int64) *rest_errors.RestErr {
	usr := &users.User{Id: userId}
	if err := usr.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *userService) SearchUser(status string) (users.Users, *rest_errors.RestErr) {
	d := &users.User{}
	return d.FindByStatus(status)
}

func (s *userService) LoginUser(login users.Login) (*users.User, *rest_errors.RestErr) {
	d := &users.User{
		Email: login.Email,
		Password: crypto.GetMd5(login.Password),
		Status: users.StatusActive,
	}
	if err := d.FindByLogin(); err != nil {
		return nil, err
	}
	return d, nil
}
