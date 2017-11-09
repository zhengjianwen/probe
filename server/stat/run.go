package stat

import (
	"github.com/go-xorm/xorm"
	"github.com/rongyungo/probe/server/apm"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/util/sql"
	"log"
	"time"
)

var Orm *xorm.Engine

func InitMySQL(cfg *sql.DatabaseConfig) (err error) {
	Orm, err = sql.InitMySQL(cfg)
	return sync()
}

func Start() {
	tk1 := time.NewTicker(time.Second * time.Duration(20))
	for {
		select {
		case <-tk1.C:
			CalculateTaskAvaliablilty()
		case task := <-reduceCh:
			if err := ReduceScheduleTask(task); err != nil {
				log.Printf("[stat] reduce schedule task err %v\n", err)
			}
		}
	}
}

func sync() error {
	return Orm.Sync2(
		new(types.TaskStat),
		new(types.TaskSchedule),
	)
}

func CalculateTaskAvaliablilty() {
	var l []types.TaskSchedule
	if err := Orm.Distinct("task_type", "task_id").Find(&l); err != nil {
		log.Printf("<<<<<<<<<<<<<<<distinct task err %v>>>>>>>>>>>>>> ", err)
		return
	}

	for _, ts := range l {
		var taskList types.TaskScheduleList
		err := Orm.Where("task_type = ? AND task_id = ? AND (UNIX_TIMESTAMP()-schedule_time) <= 3 * period_sec", ts.TaskType, ts.TaskId).
			Find(&taskList)
		if err != nil {
			log.Printf("<<<<<<<<<<<<<<<handler one task err %v>>>>>>>>>>>>>> ", err)
			continue
		}

		if task := taskList.ReturnFinishedTask(); task != nil {
			total := float64(task.SuccessN + task.ErrorN)
			av := int(float64(task.SuccessN) / total * 100)
			delay := int(float64(task.DelaySum) / total)
			err := apm.PushHttpStat(task.TaskId, task.OrgId, av, delay, int(task.PeriodSec))
			if err != nil {
				log.Printf("[stat] push http stat err %v\n", err)
				continue
			}

			reduceCh <- task
		}
	}
}
