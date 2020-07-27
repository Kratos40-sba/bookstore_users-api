package mysql_utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kratos40-sba/bookstore_users-api/utils/errors"
	"strings"
)
const(
	NoUserFound = " no rows in result set"
)
func ParseErr (err error) *errors.RestErr{
	sqlErr , ok := err.(*mysql.MySQLError)
	if !ok{
		if strings.Contains(err.Error(),NoUserFound){
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerErrors("error parsing database response")
	}
	switch sqlErr.Number {
	case 1062 :
		return errors.NewBadRequest("invalid data")
		}
		return errors.NewInternalServerErrors("error processing request")
}
func HandleInternalServerErr(err error) *errors.RestErr{
	if err !=nil {
		return errors.NewInternalServerErrors(err.Error())
	}
	return nil
}