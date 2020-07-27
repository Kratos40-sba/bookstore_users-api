package app

import (
	"github.com/kratos40-sba/bookstore_users-api/controllers/user"
)

func mapUrls(){
// user
router.POST("/users", user.CreateUser)
router.GET("/users/:id", user.GetUser)
router.GET("/users", user.FindUser)
}
