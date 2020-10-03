package taskrunner

import (
	"awesomeProject3/scheduler/dbops"
	"errors"
)

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3)
	if err != nil {
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished.")
	}
	for _, id := range res {
		dc <- id
	}
	return nil
}
