package dbops

import (
	"awesomeProject3/video/api/defs"
	"awesomeProject3/video/api/utils"
	"database/sql"
	"log"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users value(login_name, pwd) values(?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtIns, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		return "", err
	}
	var pwd string
	stmtIns.QueryRow(loginName).Scan(&pwd)
	stmtIns.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")
	if err != nil {
		log.Printf("ddd")
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`insert into video_info (id, author_id, name, display_ctime)values(?,?,?,?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}

	defer stmtIns.Close()
	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtIns, err := dbConn.Prepare("select author_id, name, display_ctime from video_info where vid = ?")

	var aid int
	var dct string
	var name string

	err = stmtIns.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err != sql.ErrNoRows {
		return nil, nil
	}
	stmtIns.Close()
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}
	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtIns, err := dbConn.Prepare("delete from video_info where vid = ?")
	if err != nil {
		log.Printf("ddd")
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		return err
	}
	stmtIns.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare(`insert into comments (id, video_id, author_id, content)values(?,?,?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtIns, err := dbConn.Prepare(`select comments.id, users.Login_name, comments.content from comments inner join users on comments.author_id = users.id where comments.video_id = ? and comments.time > FROM_UNIXTIME(?) and commtents.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err = stmtIns.Query(vid, from,to)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err != sql.ErrNoRows {
		return nil, nil
	}
	stmtIns.Close()

	return res, nil
}
}