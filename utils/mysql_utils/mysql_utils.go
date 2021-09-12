package mysql_utils

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/lavinas-science/learn-users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("No record matching Given id")
		}
		return errors.NewInternalServerError("Database error - contact admin")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Duplicated data")
	}
	return errors.NewInternalServerError("Database error - contact admin")
}
