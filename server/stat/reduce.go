package stat

import (
	"strconv"
	"fmt"
	"errors"
	"log"
	"github.com/rongyungo/probe/server/master/types"
)

// reduce is used to merge some task_schedule into one raw
// as table task_schedule Space complexity is task * scheduleTime

func Reduce() error {
	tasks, err := distinctTaskSchedule()
	if err != nil {
		return err
	}

	log.Printf("<<<<<<<<<<<<%#v>>>>>>>>>>\n ", tasks)

	for _, task := range tasks {
		//当在task schedule表中, 某个task存在20条以上的调度信息时，需要清除，只留最新的20条
		if task.TaskScheduleCount > 20 {
			//查询需要清除的task schedule 的 调度id(scheduleTime)
			scheduleTime, err := queryTaskScheduleTime(task.TaskType, task.TaskId)
			if err != nil {
				log.Printf("worker stat query schedule time err %v\n", err)
				continue
			}

			//统计自scheduleTime以后该task的所有成功和失败总和
			successN, errorN, err := sumTaskInfo(task.TaskType, task.TaskId, int64(scheduleTime))
			if err != nil {
				log.Printf("sum task schedual success/error err %v", err)
				continue
			}

			err = appendTaskStat(task.TaskType, task.TaskId, successN, errorN)
			if err != nil {
				log.Printf("append task(%d) schedule result err %v\n", task.TaskId,  err)
				continue
			}

			n, err := delOldSchedules(task.TaskType, task.TaskId, int64(scheduleTime))
			if err != nil {
				log.Printf("delete task(%d) %d  schedule rows", task.TaskId, n)
				continue
			}
		}
	}

	return nil
}

type taskCount struct {
	TaskType 		  int64 `xorm:"task_type"`
	TaskId            int64 `xorm:"task_id"`
	TaskScheduleCount int64 `xorm:"task_schedule_count"`
}

func distinctTaskSchedule() ([]taskCount, error) {
	sql := `select task_type, task_id, count(schedule_time) as task_schedule_count
		from task_schedule group by task_id, task_type order by task_schedule_count DESC`
	res, err := Orm.Query(sql)
	if err != nil {
		return nil, err
	}

	var l []taskCount
	for _, rawMap := range res{
		task_type, err := readIntColumn(rawMap, "task_type")
		if err != nil {
			continue
		}
		task_id, err := readIntColumn(rawMap, "task_id")
		if err != nil {
			continue
		}
		task_schedule_count, err := readIntColumn(rawMap, "task_schedule_count")
		if err != nil {
			continue
		}

		l = append(l, taskCount{
			TaskType: 		   int64(task_type),
			TaskId:  		   int64(task_id),
			TaskScheduleCount: int64(task_schedule_count),
		})
	}
	return l, nil
}

func queryTaskScheduleTime(ttp, tid int64) (int, error) {
	sql := fmt.Sprintf(`select min(a.schedule_time) as schedule_time from (
		select schedule_time from task_schedule where task_type = %d AND task_id = %d order by schedule_time DESC limit 20
	) as a`, ttp, tid)
	res, err := Orm.Query(sql)
	if err != nil {
		return 0, err
	}

	if len(res) > 0 {
		return readIntColumn(res[0], "schedule_time")
	}

	return 0, errors.New("not found")
}

func sumTaskInfo(ttp, tid, scheduleTime int64) (int64, int64, error) {
	sql := fmt.Sprintf(`select sum(success_n) success_n, sum(error_n) error_n
		from task_schedule where task_type = %d AND task_id = %d and schedule_time <= %d`, ttp, tid, scheduleTime)
	res, err := Orm.Query(sql)
	if err != nil {
		return 0, 0, err
	}

	if len(res) == 0 {
		return 0, 0, errors.New("sum task success and error fail")
	}

	successN, err := readIntColumn(res[0], "success_n")
	if err != nil {
		return 0, 0, err
	}
	errorN, err := readIntColumn(res[0], "error_n")
	if err != nil {
		return 0, 0, err
	}

	return int64(successN), int64(errorN), nil
}

func appendTaskStat(ttp, tid, successN, errorN int64) error {
	ts := types.TaskStat{
		TaskType: 	ttp,
		TaskId: 	tid,
		SuccessN: 	successN,
		ErrorN: 	errorN,
	}
	n, err := Orm.Table(ts).Where("task_type = ? AND task_id = ?", ttp, tid).Count()
	if err != nil {
		return err
	}

	if n == 0 {
		_, err = Orm.InsertOne(ts)
	} else {
		_, err= Orm.Exec("UPDATE task_stat SET success_n = success_n + ? AND error_n = error_n + ? WHERE task_type = ? AND task_id = ?",
			successN, errorN, ttp, tid)
	}
	return err
}

func delOldSchedules(ttp, tid, sid int64) (int64, error) {
	res, err := Orm.Exec("delete from task_schedule where task_type = ? AND task_id = ? and schedule_time <= ? ", ttp, tid, sid)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

var GetColumnNotFoundErr = func(column string) error {
	return fmt.Errorf("column %d not found", column)
}

func readIntColumn(m map[string][]byte, column string) (int, error) {
	if len(m) == 0 {
		return 0, GetColumnNotFoundErr(column)
	}

	cData, ok := m[column]
	if !ok {
		return 0, GetColumnNotFoundErr(column)
	}

	return strconv.Atoi(string(cData))
}
