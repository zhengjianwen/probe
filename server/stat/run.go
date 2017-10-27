package stat

import (
	"github.com/rongyungo/probe/util/sql"
	"github.com/go-xorm/xorm"
	"github.com/rongyungo/probe/server/master/types"
	"time"
	//"github.com/rongyungo/probe/server/apm"
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

	//当存在2组以上的task result没有处理时， 处理最old的item
	if len(l) >= 2 {
		log.Printf("---------> get result %#v\n <---------------", l[0])
		_, err := Orm.Where("worker_id = ? AND task_id = ? AND schedule_time = ?",
			l[0].WorkerId, l[0].TaskId, l[0].ScheduleTime).Cols("if_stat").Update(types.TaskSchedule{
				IfStat: true,
		})
		if err != nil {
			log.Printf("[stat] update task schedule result stat err %v\n", err)
			return
		}

		//apm.PushHttpStat(l[0].ta)

	}
}

