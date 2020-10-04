package dbops

import "log"

//api->videoid->mysql
//dispatcher->mysql->videoid->datachannel
//executor->datachannel->videoid->delete videos
func AddVideoDeletionRecord(video string) error {
	stmtOut, err := dbConn.Prepare("insert into video_del_rec (video_id) values(?)")

	if err != nil {
		return err
	}
	_, err = stmtOut.Exec(video)

	if err != nil {
		log.Printf("Add VideoDeletionRecord err: %v", err)
		return err
	}
	defer stmtOut.Close()
	return nil
}

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_id from video_del_rec limit ?")

	var ids []string
	if err != nil {
		return ids, err
	}
	rows, err := stmtOut.Query(count)

	if err != nil {
		log.Printf("Query VideoDeletionRecord err: %v", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("delete from video_del_rec where video_id = ?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecord err: %v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
