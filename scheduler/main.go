package main

import (
	"awesomeProject3/scheduler/taskrunner"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	router *httprouter.Router
}

func (m *middleWareHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.router.ServeHTTP(writer, request)
}

func NewMiddleWareHandler(router *httprouter.Router) *middleWareHandler {
	return &middleWareHandler{router: router}
}

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)
	return router
}

func main() {
	go taskrunner.Start()
	http.ListenAndServe(":9001", NewMiddleWareHandler(RegisterHandler()))
}
