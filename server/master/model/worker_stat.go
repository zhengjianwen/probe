package model

import (
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/server/master/types"
	"fmt"
	"time"
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

//taskType to taskId to workerId to DelayMs mapping
var TaskSnapShotMapping = map[string]map[int64]map[int64]struct{
	SnapShotTimeStamp int64
	DelayMs           int64
}{
	"HTTP": make(map[int64]map[int64]struct{
	SnapShotTimeStamp int64
	DelayMs           int64}),
	"DNS": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
	"PING": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
	"FTP": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
	"TCP": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
	"UDP": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
	"TRACE_ROUTE": make(map[int64]map[int64]struct{
		SnapShotTimeStamp int64
		DelayMs           int64}),
}

func CoverSnapShotM(tp string, tid, wid int64, delayMs int64) {
	if _, ok := TaskSnapShotMapping[tp]; !ok {
		return
	}

	if _, ok := TaskSnapShotMapping[tp][tid]; !ok {
		TaskSnapShotMapping[tp][tid] = make(map[int64]struct{
			SnapShotTimeStamp int64
			DelayMs int64
		})
	}

  	TaskSnapShotMapping[tp][tid][wid] = struct{
		SnapShotTimeStamp int64
		DelayMs int64
  	}{
		SnapShotTimeStamp : time.Now().Unix(),
		DelayMs: delayMs,
  	}
}