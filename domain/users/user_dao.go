package users

import (
	"fmt"

	"github.com/lavinas-science/learn-users-api/utils/dates"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	r := userDB[user.Id]
	if r == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Id = r.Id
	user.FirstName = r.FirstName
	user.LastName = r.LastName
	user.Email = r.Email
	user.DateCreated = r.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	db := userDB[user.Id]
	if db != nil {
		if db.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %d already exists", user.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	user.DateCreated = dates.GetNowString()

	userDB[user.Id] = user
	return nil
}
