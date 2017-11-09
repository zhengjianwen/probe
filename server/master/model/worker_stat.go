package model

import (
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
	"sync"
	"time"
)

func SyncTaskResult(res *pb.TaskResult) error {
	//exist, err := Orm.Where("task_type = ? AND task_id = ? AND schedule_time = ?", int(res.GetType()), res.TaskId, res.ScheduleTime).Get(&wks)
	//if err != nil {
	//	return err
	//}
	//
	//if !exist {
	//	return CreateTaskResult(res)
	//}
	var setSql = "success_n = success_n + 1"
	if !res.Success {
		setSql = "error_n = error_n +1"
	}
	sql := fmt.Sprintf("UPDATE task_schedule SET delay_sum = delay_sum + ?, %s WHERE task_type = ? AND task_id = ? AND schedule_time = ?", setSql)

	_, err := Orm.Exec(sql, res.DelayMs, int(res.GetType()), res.TaskId, res.ScheduleTime)
	return err
}

var l *sync.Mutex = new(sync.Mutex)
var HttpSnapShotMapping = map[int64]map[int64]struct {
	SnapShotTimeStamp int64
	DelayMs           int64
}{}

func CoverSnapShotM(tp string, tid, wid int64, delayMs int64) {
	now := time.Now().Unix()
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
				SnapShotTimeStamp: now,
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

	//WorkTimeStampMapping[wid] = now
}

//var wl *sync.RWMutex = new(sync.RWMutex)
//// worker id to latest working timestamp to validate if the worker alive
//var WorkTimeStampMapping = map[int64]int64{}
//
//func getWorkingWorker() int {
//	wl.RLock()
//	defer wl.RUnlock()
//	var total int
//
//	total, now := 0, time.Now()
//	for _, ts := range WorkTimeStampMapping {
//		if  now.Sub(time.Unix(ts, 0)) <= time.Minute * 5 {
//			total ++
//		}
//	}
//
//	return total
//}
