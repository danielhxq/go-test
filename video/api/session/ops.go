package session

import (
	"awesomeProject3/video/api/dbops"
	"awesomeProject3/video/api/defs"
	"awesomeProject3/video/api/utils"
	"sync"
	"time"
)

var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func LoadSessionFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}
	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func deleteExpiredFromDB(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := time.Now().UnixNano() / 10000000
	ttl := ct + 30*60*1000

	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)

	return id
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	if ok {
		ct := time.Now().UnixNano() / 1000000
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredFromDB(sid)
			return "", true
		}

		return ss.(*defs.SimpleSession).Username, false
	}
	return "", true
}
