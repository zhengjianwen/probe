package model

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func HandleTaskResult(wid int64, r *pb.TaskResult) error {
	if r.Type == pb.TaskType_HTTP {
		return SyncTackScheduleResult(wid, r)
	}
	//_, err := Orm.Insert(r)
	return nil
}
