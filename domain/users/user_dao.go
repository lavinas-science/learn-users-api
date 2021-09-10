package users

import (
	"github.com/lavinas-science/learn-users-api/datasources/mysql/users_db"
	"github.com/lavinas-science/learn-users-api/utils/dates"
	"github.com/lavinas-science/learn-users-api/utils/errors"
	"github.com/lavinas-science/learn-users-api/utils/mysql_utils"
)


const (
	queryInsertUser = "insert into users(first_name, last_name, email, date_created) values (?, ?, ?, ?);"
	queryGetUser = "select id, first_name, last_name, email, date_created from users where id = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryGetUser)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	res := stmt.QueryRow(user.Id)
	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	user.DateCreated = dates.GetNowString()
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	// can replace all above by this
	// res, saveRrr := users_db.Db.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)	
	}
	user.Id = userId
	return nil
}
