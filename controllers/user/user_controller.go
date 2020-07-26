package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kratos40-sba/bookstore_users-api/domain/users"
	"github.com/kratos40-sba/bookstore_users-api/services"
	"github.com/kratos40-sba/bookstore_users-api/utils/errors"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	id , userErr := strconv.ParseInt(c.Param("id"),10,64)
	if userErr !=nil{
		err := errors.NewBadRequest("invalid user id")
		c.JSON(err.Status,err)
		return
	}
	result, getErr := services.GetUser(id)
	if getErr != nil {
		// handle user getting error
		c.JSON(getErr.Status, getErr)
		return
		//log.Fatal(saveErr) }

	//c.String(http.StatusNotImplemented,"not yet implemented \n")
    }
	c.JSON(http.StatusOK, result)
}
func CreateUser(c *gin.Context) {
	var user users.User
	/*body , err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body , &user) ; err !=nil{
		log.Fatal(err)
	}*/
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequest("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
		//log.Fatal(err)
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
		//log.Fatal(saveErr) }


	}
	c.JSON(http.StatusCreated, result)
}
func FindUser(c *gin.Context){
	c.String(http.StatusNotImplemented,"not yet implemented \n")
}
