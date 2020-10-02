package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	http.ListenAndServe(":8000", RegisterHandlers())
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)
	router.GET("/user/:user_name/videos", CreateUser)
	router.GET("/user/:user_name/videos/:vid-id", CreateUser)
	router.DELETE("/user/:user_name/videos/:vid-id", CreateUser)

	router.GET("/videos/:vid-id/comments", CreateUser)
	router.POST("/videos/:vid-id/comments", CreateUser)
	router.DELETE("/videos/:vid-id/comment/:comment-id", CreateUser)
	return router
}
