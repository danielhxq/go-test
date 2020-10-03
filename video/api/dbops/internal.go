package dbops

import (
	"awesomeProject3/video/api/defs"
	"database/sql"
	"log"
	"strconv"
	"sync"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlstr := strconv.FormatInt(ttl, 10)

	stmtIns, err := dbConn.Prepare("insert into sessions (session_id, TTL, login_name) values(?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(sid, ttlstr, uname)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	ss := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("select TTL, login_name from sessions where session_id = ?")
	if err != nil {
		return nil, err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	//var ttlint int64
	if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		ss.TTL = res
		ss.Username = uname
	} else {
		return nil, err
	}
	defer stmtOut.Close()
	return ss, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	m := &sync.Map{}
	stmtIns, err := dbConn.Prepare(`select TTL, login_name from sessions`)
	if err != nil {
		return nil, err
	}

	rows, err := stmtIns.Query()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err != sql.ErrNoRows {
		return nil, nil
	}
	defer stmtIns.Close()

	for rows.Next() {
		var id, ttlstr, login_name string
		if err := rows.Scan(&id, &ttlstr, &login_name); err != nil {
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			ss := &defs.SimpleSession{Username: login_name, TTL: ttl}
			m.Store(id, ss)
		}
	}
	return m, nil
}

func DeleteSession(sid string) error {
	stmtIns, err := dbConn.Prepare("delete from sessions where session_id = ?")
	if err != nil {
		log.Printf("ddd")
		return err
	}
	_, err = stmtIns.Exec(sid)
	if err != nil {
		return err
	}
	stmtIns.Close()
	return nil
}
