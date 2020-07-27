package users

import (
	"github.com/kratos40-sba/bookstore_users-api/datasources/mysql/user_db"
	"github.com/kratos40-sba/bookstore_users-api/utils/date_utils"
	"github.com/kratos40-sba/bookstore_users-api/utils/errors"
	"github.com/kratos40-sba/bookstore_users-api/utils/mysql_utils"
)
const (
	queryInsertUser = "INSERT INTO user(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM user WHERE id=?;"

)
// only here we interacte with the DataBase
func (user *User) Get() *errors.RestErr{
	stmt , err := user_db.Client.Prepare(queryGetUser)
	mysql_utils.HandleInternalServerErr(err)
	defer stmt.Close()
	res := stmt.QueryRow(user.Id)
	if getErr := res.Scan(&user.Id,&user.FirstName,&user.LastName,&user.Email,&user.DateCreated) ; err != nil {
	return mysql_utils.ParseErr(getErr)
	}
	return  nil
}
func (user *User)Save()*errors.RestErr{
	stmt , err := user_db.Client.Prepare(queryInsertUser)
	mysql_utils.HandleInternalServerErr(err)
	// very importatnt to close
	defer stmt.Close()
	user.DateCreated = date_utils.GetNowString();
	insertRes , saveErr := stmt.Exec(user.FirstName,user.LastName,user.Email,user.DateCreated)
	if saveErr != nil {
	return mysql_utils.ParseErr(saveErr)
	}
	userId , err := insertRes.LastInsertId()
	if err !=nil{
		return mysql_utils.ParseErr(saveErr)
	}
	user.Id = userId
	return nil
}
