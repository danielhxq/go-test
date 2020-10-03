package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func main() {
	r := NewMiddleWareHandler(RegisterHandlers())
	http.ListenAndServe(":8000", r)
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
