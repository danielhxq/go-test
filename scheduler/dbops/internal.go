package dbops

import "log"

//api->videoid->mysql
//dispatcher->mysql->videoid->datachannel
//executor->datachannel->videoid->delete videos

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
