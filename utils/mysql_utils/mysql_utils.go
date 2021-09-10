package mysql_utils

import (
	"fmt"
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
			return errors.NewNotFoundError("No record matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("Error parsing database response: %s", err.Error()))
	}
	switch sqlErr.Number {
		case 1062: 
			return errors.NewBadRequestError("Duplicated data")
	}
	return errors.NewInternalServerError(fmt.Sprintf("Error processing request: %s", err.Error()))
}
