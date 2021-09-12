package users

import (
	"fmt"

	"github.com/lavinas-science/learn-users-api/datasources/mysql/users_db"
	"github.com/lavinas-science/learn-users-api/utils/errors"
	"github.com/lavinas-science/learn-users-api/utils/mysql_utils"
	"github.com/lavinas-science/learn-users-api/utils/logger"
)

const (
	queryInsertUser   = "insert into users(first_name, last_name, email, date_created, status, password) values (?, ?, ?, ?, ?, ?);"
	queryGetUser      = "select id, first_name, last_name, email, date_created, status from users where id = ?;"
	queryUpdateUser   = "update users set first_name = ?, last_name = ?, email = ? where id = ?;"
	queryDeleteUser   = "delete from users where id = ?;"
	queryFindByStatus = "select id, first_name, last_name, email, date_created, status from users where status = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user stat", err)
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	res := stmt.QueryRow(user.Id)
	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to scan get user stat", err)
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user stat", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	// can replace all above by this
	// res, saveRrr := users_db.Db.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		logger.Error("error when trying to exec save user stat", err)
		return mysql_utils.ParseError(err)
	}
	userId, err := res.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user stat", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, errUp := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if errUp != nil {
		logger.Error("error when trying to exec update user stat", errUp)
		return mysql_utils.ParseError(errUp)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user stat", err)
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, errUp := stmt.Exec(user.Id)
	if errUp != nil {
		logger.Error("error when trying to exec delete user stat", errUp)
		return mysql_utils.ParseError(errUp)
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Db.Prepare(queryFindByStatus)
	if err != nil {
		logger.Error("error when trying to prepare users find by status stat", err)
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rq, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to query users find by status stat", err)
		return nil, mysql_utils.ParseError(err)
	}
	defer rq.Close()
	r := make([]User, 0)
	for rq.Next() {
		var us User
		if err := rq.Scan(&us.Id, &us.FirstName, &us.LastName,
			&us.Email, &us.DateCreated, &us.Status); err != nil {
			logger.Error("error when trying to scan users find by status stat", err)
			return nil, mysql_utils.ParseError(err)
		}
		r = append(r, us)
	}
	if len(r) == 0 {
		// logger.Info("No user matching status when user find by status")
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return r, nil
}
