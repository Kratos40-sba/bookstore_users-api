package users

import (
	"fmt"
	"github.com/kratos40-sba/bookstore_users-api/utils/errors"
)

// only here we interacte with the DataBase
var (
	userDB = make(map[int64]*User)
)
func (user *User) Get() *errors.RestErr{
	userFound := userDB[user.Id]
	if userFound == nil {
       return errors.NewNotFoundError(fmt.Sprintf("user of the id %d not found \n",user.Id ))
	}
	user.Id = userFound.Id
	user.FirstName = userFound.FirstName
	user.LastName = userFound.LastName
	user.Email = userFound.Email
	user.DateCreated = userFound.DateCreated
	return  nil
}
func (user *User)Save()*errors.RestErr{
	currentUser := userDB[user.Id]
	if currentUser != nil {
		return errors.NewBadRequest(fmt.Sprintf("user of the id %d already exists",user.Id))
	}
	userDB[user.Id] = user
	return nil
}
