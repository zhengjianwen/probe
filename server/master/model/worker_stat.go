package model

import (
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/server/master/types"
	"fmt"
	"time"
	"sync"
)

func SyncTaskResult(res *pb.TaskResult) error {
	var wks types.TaskSchedule
	exist, err := Orm.Where("task_id = ? AND schedule_time = ?", res.TaskId, res.ScheduleTime).Get(&wks)
	if err != nil {
		return err
	}

	if !exist {
		return CreateTaskResult(res)
	}

	sql := "UPDATE task_schedule SET delay_sum = delay_sum + ?, %s WHERE task_id = ? AND schedule_time = ? AND period_sec = ?"
	var setSql = "success_n = success_n + 1"
	if !res.Success {
		setSql = "error_n = error_n +1"
	}

	_, err = Orm.Exec(fmt.Sprintf(sql, setSql), res.DelayMs, res.TaskId, res.ScheduleTime, res.PeriodSec)
	return err
}

func CreateTaskResult(res *pb.TaskResult) error {
	session := Orm.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return err
	}

	ts := types.TaskSchedule{
		TaskType: 		int64(int32(res.Type)),
		TaskId: 		res.TaskId,
		DelaySum:  		res.DelayMs,
		ScheduleTime: 	res.ScheduleTime,
		PeriodSec: 		int32(res.PeriodSec),
	}

	if res.Success {
		ts.SuccessN = 1
	} else {
		ts.ErrorN = 1
	}

	_, err = session.InsertOne(ts)
	return session.Commit()
}

var l *sync.Mutex = new(sync.Mutex)
var HttpSnapShotMapping = map[int64]map[int64]struct{
	SnapShotTimeStamp int64
	DelayMs           int64
}{}

func CoverSnapShotM(tp string, tid, wid int64, delayMs int64) {
	switch tp {
	case "HTTP":
		l.Lock()
		defer l.Unlock()
		if wm, ok := HttpSnapShotMapping[tid]; !ok {
			HttpSnapShotMapping[tid] = make(map[int64]struct {
				SnapShotTimeStamp int64
				DelayMs           int64
			})
		} else {
			wm[wid] = struct {
				SnapShotTimeStamp int64
				DelayMs           int64
			}{
				SnapShotTimeStamp: time.Now().Unix(),
				DelayMs:           delayMs,
			}
		}

	case "DNS":
	case "PING":
	case "FTP":
	case "TCP":
	case "UDP":
	case "TRACE_ROUTE":
		return
	}

}