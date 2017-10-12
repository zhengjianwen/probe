package scheduler

import (
	"github.com/go-xorm/xorm"

	pb "github.com/rongyungo/probe/server/proto"
	errutil "github.com/rongyungo/probe/util/errors"
	sqlutil "github.com/rongyungo/probe/util/sql"
	"strings"

	"log"
	"time"
)

type ScheduleManager struct {
	TaskType    pb.TaskType
	PeriodSec   uint8
	DbConfig    *sqlutil.DatabaseConfig
	Db          *xorm.Engine
	taskManager *taskManager
}

func NewSchedulerManager(tp string, period uint8, c *sqlutil.DatabaseConfig) (*ScheduleManager, error) {
	tpId, ok := pb.TaskType_value[strings.ToUpper(tp)]
	if !ok {
		return nil, errutil.ErrUnSupportTaskType
	}

	if period < 60 {
		return nil, errutil.ErrTaskPeriodTooLess
	}

	engine, err := sqlutil.InitMySQL(c)
	if err != nil {
		return nil, err
	}

	return &ScheduleManager{
		TaskType:    pb.TaskType(tpId),
		PeriodSec:   period,
		DbConfig:    c,
		Db:          engine,
		taskManager: NewTaskManager(),
	}, nil
}

func (m *ScheduleManager) Start() error {
	go m.Run()

	tks, err := m.GetAllTasks()
	if err != nil {
		return err
	}
	m.taskManager.SyncTask(tks)
	m.taskManager.StatTasks()
	return nil
}

func (m *ScheduleManager) Run() {
	go m.CorrectScheduleTime()
	ctk := time.NewTicker(time.Second * time.Duration(m.PeriodSec)) //scheduler time correct ticker
	tk := time.NewTicker(time.Second * time.Duration(5))

	for {
		select {
		case <-tk.C:
			tasks, err := m.GetScheduleTasks()
			if err != nil {
				log.Printf("scheduler get to schedule tasks err %v\n", err)
				continue
			}

			tasks = m.taskManager.ReduceReplicatedTask(tasks)
			if len(tasks) > 0 {
				log.Printf("scheduler[%s]: query prepare schedule tasks: %d\n", m.TaskType.String(), len(tasks))
			}

			m.Schedule(nil, tasks)

		case <-ctk.C:
			go m.CorrectScheduleTime()
		}
	}
}
