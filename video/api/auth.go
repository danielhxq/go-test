package main

import (
	"awesomeProject3/video/api/dbops"
	"awesomeProject3/video/api/defs"
	"awesomeProject3/video/api/session"
	"net/http"
)

var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return false
	}

	uname, ok := session.IsSessionExpired(sid)
	if ok {
		return false
	}
	r.Header.Add(HEADER_FIELD_UNAME, uname)
	return true
}

func ValidateUser(w http.ResponseWriter, r *http.Request) bool {
	uname := r.Header.Get(HEADER_FIELD_UNAME)
	if len(uname) == 0 {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	_, err := dbops.RetrieveSession(sid)
	if err != nil {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	if _, ok := session.IsSessionExpired(sid); !ok {
		sendErrorResponse(w, defs.ErrorNotAuthUser)
		return false
	}
	return true
}
