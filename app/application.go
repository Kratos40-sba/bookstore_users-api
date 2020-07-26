package app

import "github.com/gin-gonic/gin"

var (
	// router that we using
	router = gin.Default()
)
func StartApplication (){
	mapUrls()
	router.Run()
	// this is the only point where we using an http server


}
