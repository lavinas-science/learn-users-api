package users

import "encoding/json"

type PublicUser struct {
	Id int64 `json:"id"`
	// FirstName   string `json:"first_name"`
	// LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var pu PrivateUser
	json.Unmarshal(userJson, &pu)
	return pu
}

func (users Users) Marshall(isPublic bool) []interface{} {
	r := make([]interface{}, len(users))
	for i, u := range users {
		r[i] = u.Marshall(isPublic)
	}
	return r
}
