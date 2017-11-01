package model

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func HandleTaskResult(r *pb.TaskResult) error {
	if r.Type == pb.TaskType_HTTP {
		return SyncTackScheduleResult(r)
	}
	return nil
}
