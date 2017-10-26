package model

import (
	"github.com/rongyungo/probe/server/master/types"
	"fmt"
)

func SyncTackScheduleResult(wid, tid, scheduleT int64, ok bool) error {
	var wks types.TaskSchedule
	exist, err := Orm.Where("worker_id = ? AND task_id = ? AND schedule_time = ?",
		wid, tid, scheduleT).Get(&wks)
	if err != nil {
		return err
	}

	if !exist {
		return CreateTaskSchedule(wid, tid, scheduleT, ok )
	}

	sql := "UPDATE task_schedule SET %s WHERE worker_id = ? AND task_id = ? AND schedule_time = ?"
	var setSql = "success_n = success_n + 1"
	if !ok {
		setSql = "error_n = error_n +1"
	}

	_, err = Orm.Exec(fmt.Sprintf(sql, setSql), wid, tid, scheduleT)
	return err
}

func CreateTaskSchedule(wid, tid, scheduleT int64, ok bool) error {
	session := Orm.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return err
	}

	ts := types.TaskSchedule{
		WorkerId: wid,
		TaskId: tid,
		ScheduleTime: scheduleT,
	}

	if ok {
		ts.SuccessN = 1
	} else {
		ts.ErrorN = 1
	}

	_, err = session.InsertOne(ts)
	return session.Commit()
}
//func StatTaskAvailabilityInHours(wid, tid, h int64) (float64, error) {
//	var l []types.WorkerStatHour
//	nowHour := time.Now().Truncate(time.Hour * time.Duration(h)).Format("2006-01-02 15")
//	if err := Orm.Where("worker_id = ? AND task_id = ? AND hour >= ?",
//		wid, tid, nowHour).Find(&l); err != nil {
//			return -1, err
//	}
//
//	return statHoursAvailability(l), nil
//}

//func StatTaskAvailabilityInDays(wid, tid, h int64) (float64, error) {
//	var l []types.WorkerStatDay
//	nowHour := time.Now().Truncate(time.Hour * time.Duration(h) * 24).Format("2006-01-02 15")
//	if err := Orm.Where("worker_id = ? AND task_id = ? AND hour >= ?",
//		wid, tid, nowHour).Find(&l); err != nil {
//		return -1, err
//	}
//
//	return statDaysAvailability(l), nil
//}
//
//func statHoursAvailability(l []types.WorkerStatHour) float64 {
//	var sum_total, sum_success int64
//	for _, s := range l {
//		sum_total += s.RequestN
//		sum_success += s.SuccessN
//	}
//
//	if sum_total == 0 {
//		return 0
//	} else {
//		return float64(sum_success) / float64(sum_total)
//	}
//}
//
//func statDaysAvailability(l []types.WorkerStatDay) float64 {
//	var sum_total, sum_success int64
//	for _, s := range l {
//		sum_total += s.RequestN
//		sum_success += s.SuccessN
//	}
//
//	if sum_total == 0 {
//		return 0
//	} else {
//		return float64(sum_success) / float64(sum_total)
//	}
//}