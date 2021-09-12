package users

import (
	"fmt"

	"github.com/lavinas-science/learn-users-api/datasources/mysql/users_db"
	"github.com/lavinas-science/learn-users-api/utils/errors"
	"github.com/lavinas-science/learn-users-api/utils/mysql_utils"
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
		return mysql_utils.ParseError(err)
	}
	defer stmt.Close()
	res := stmt.QueryRow(user.Id)
	if err := res.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
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
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, errUp := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if errUp != nil {
		return mysql_utils.ParseError(errUp)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, errUp := stmt.Exec(user.Id)
	if errUp != nil {
		return mysql_utils.ParseError(errUp)
	}
	return nil
}

func (user *User) FindByStatus(status string) (Users, *errors.RestErr) {
	stmt, err := users_db.Db.Prepare(queryFindByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rq, err := stmt.Query(status)
	if err != nil {
		return nil, mysql_utils.ParseError(err)
	}
	defer rq.Close()
	r := make([]User, 0)
	for rq.Next() {
		var us User
		if err := rq.Scan(&us.Id, &us.FirstName, &us.LastName,
			&us.Email, &us.DateCreated, &us.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		r = append(r, us)
	}
	if len(r) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return r, nil
}
