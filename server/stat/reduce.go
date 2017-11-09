package stat

import (
	"fmt"
	"log"
	"github.com/rongyungo/probe/server/master/types"
)

// reduce is used to merge some task_schedule into one raw
// as table task_schedule Space complexity is task * scheduleTime

var reduceCh chan *types.TaskSchedule = make(chan *types.TaskSchedule, 5000)

func ReduceScheduleTask(task *types.TaskSchedule) error {
	err := appendTaskStat(task)
	if err != nil {
		log.Printf("append task(%d) schedule result err %v\n", task.TaskId,  err)
		return err
	}
	_, err = removeScheduleTask(task)
	return err
}

func appendTaskStat(task *types.TaskSchedule) error {
	ts := types.TaskStat{
		TaskType: 	task.TaskType,
		TaskId: 	task.TaskId,
		SuccessN: 	task.SuccessN,
		ErrorN: 	task.ErrorN,
		DelaySum:   task.DelaySum,
	}
	n, err := Orm.Table(ts).Where("task_type = ? AND task_id = ?", task.TaskType, task.TaskId).Count()
	if err != nil {
		return err
	}

	if n == 0 {
		_, err = Orm.InsertOne(ts)
	} else {
		sql := fmt.Sprintf("UPDATE task_stat SET success_n = success_n + %d, error_n = error_n + %d, delay_sum = delay_sum + %d WHERE task_type = %d AND task_id = %d",
			task.SuccessN, task.ErrorN, task.DelaySum, task.TaskType, task.TaskId)
		_, err= Orm.Exec(sql)
	}
	return err
}

func removeScheduleTask(task *types.TaskSchedule) (int64, error) {
	return Orm.Where("task_type = ? AND task_id = ? and schedule_time = ?",
		task.TaskType, task.TaskId, task.ScheduleTime).Delete(task)
}
