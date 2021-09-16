package services

import (
	"github.com/lavinas-science/learn-users-api/domain/users"
	"github.com/lavinas-science/learn-users-api/utils/crypto_utils"
	date_utils "github.com/lavinas-science/learn-users-api/utils/dates_utils"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

var (
	UserService userServiceInterface = &userService{}
)

type userService struct {
}

type userServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RestErr)
	GetUser(int64) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.Login) (*users.User, *errors.RestErr)
}

func (s *userService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	user.Status = users.StatusActive
	user.DateCreated = date_utils.GetNowDb()
	user.Password = crypto_utils.GetMd5(user.Password)

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) GetUser(userid int64) (*users.User, *errors.RestErr) {
	u := users.User{Id: userid}
	if err := u.Get(); err != nil {
		return nil, err
	}
	if err := u.Get(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *userService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
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

func (s *userService) DeleteUser(userId int64) *errors.RestErr {
	usr := &users.User{Id: userId}
	if err := usr.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *userService) SearchUser(status string) (users.Users, *errors.RestErr) {
	d := &users.User{}
	return d.FindByStatus(status)
}

func (s *userService) LoginUser(login users.Login) (*users.User, *errors.RestErr) {
	d := &users.User{
		Email: login.Email,
		Password: crypto_utils.GetMd5(login.Password),
		Status: users.StatusActive,
	}
	if err := d.FindByLogin(); err != nil {
		return nil, err
	}
	return d, nil
}
