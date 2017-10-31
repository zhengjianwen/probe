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
	tk := time.NewTicker(time.Second * time.Duration(20))
	for {
		select {
			case <- tk.C:
				CalculateTaskAvaliablilty()
		}
	}
}

func sync() error {
	return Orm.Sync2(
		new(types.TaskSchedule),
	)
}

func CalculateTaskAvaliablilty() {
	var l []types.TaskSchedule
	//if err := Orm.Where("if_stat = false AND (UNIX_TIMESTAMP() - schedule_time) <= 20").
	if err := Orm.Where("if_stat = false AND (UNIX_TIMESTAMP() - schedule_time) <= 60 * 10").
	OrderBy("schedule_time").Asc("schedule_time").Find(&l); err != nil {
		log.Printf("[stat] query task schedule result to calcute err %v\n", err)
	}

	if len(l) <= 1 {
		return
	}

	v := int(float64(l[0]. SuccessN)/ float64(l[0]. SuccessN + l[0].ErrorN) * 100)

	err := apm.PushHttpStat(l[0].TaskId, v, int(l[0].PeriodSec))
	if err != nil {
		log.Printf("[stat] push http stat err %v\n", err)
		return
	}

	if _, err = Orm.Where("task_id = ? AND schedule_time = ?", l[0].TaskId, l[0].ScheduleTime).
	Cols("if_stat").Update(types.TaskSchedule{IfStat: true,}); err != nil {
		log.Printf("[stat] update task schedule result stat err %v\n", err)
	}
}

