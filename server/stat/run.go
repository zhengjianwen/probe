package stat

import (
	"github.com/rongyungo/probe/util/sql"
	"github.com/go-xorm/xorm"
	"github.com/rongyungo/probe/server/master/types"
	"time"
	"github.com/rongyungo/probe/server/apm"
	"log"
)

var Orm *xorm.Engine

func InitMySQL(cfg *sql.DatabaseConfig) (err error) {
	Orm, err = sql.InitMySQL(cfg)
	return sync()
}

func Start() {
	if err := Reduce(); err != nil {
		panic(err)
	}
	tk1 := time.NewTicker(time.Second * time.Duration(20))
	tk2 := time.NewTicker(time.Minute * time.Duration(20))
	for {
		select {
			case <- tk1.C:
				CalculateTaskAvaliablilty()
			case <- tk2.C:
				Reduce()
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
	//if err := Orm.Where("if_stat = false AND (UNIX_TIMESTAMP() - schedule_time) <= 20").
	if err := Orm.Where("if_stat = false AND (UNIX_TIMESTAMP() - schedule_time) <= 60 * 10").
	OrderBy("schedule_time").Asc("schedule_time").Find(&l); err != nil {
		log.Printf("[stat] taskCount task schedule result to calcute err %av\n", err)
	}

	if len(l) <= 1 {
		return
	}

	total := float64(l[0]. SuccessN + l[0].ErrorN)
	av := int(float64(l[0]. SuccessN)/ total  * 100)
	delay := int(float64(l[0].DelaySum) / total)
	err := apm.PushHttpStat(l[0].TaskId, av, delay, int(l[0].PeriodSec))
	if err != nil {
		log.Printf("[stat] push http stat err %av\n", err)
		return
	}

	if _, err = Orm.Where("task_id = ? AND schedule_time = ?", l[0].TaskId, l[0].ScheduleTime).
	Cols("if_stat").Update(types.TaskSchedule{IfStat: true,}); err != nil {
		log.Printf("[stat] update task schedule result stat err %av\n", err)
	}
}

