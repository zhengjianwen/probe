package scheduler

import (
	"github.com/go-xorm/xorm"

	pb "github.com/rongyungo/probe/server/proto"
	errutil "github.com/rongyungo/probe/util/errors"
	sqlutil "github.com/rongyungo/probe/util/sql"
	"strings"

	"github.com/rongyungo/probe/server/master/types"
	"log"
	"reflect"
	"time"
)

type taskScheduleInfo struct {
	TaskId 		 int64
	ScheduleTime int64
}

type ScheduleManager struct {
	TaskType        pb.TaskType
	StructSliceType reflect.Type
	PeriodSec       uint8
	DbConfig        *sqlutil.DatabaseConfig
	Db              *xorm.Engine
	taskManager     *taskManager
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

	sl, ok := types.TaskTypeToStructMappings[pb.TaskType(tpId)]
	if !ok {
		return nil, errutil.ErrTaskTypeMappingNotFound
	}

	return &ScheduleManager{
		TaskType:        pb.TaskType(tpId),
		StructSliceType: reflect.TypeOf(sl),
		PeriodSec:       period,
		DbConfig:        c,
		Db:              engine,
		taskManager:     NewTaskManager(),
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
	tk := time.NewTicker(time.Second * time.Duration(5))

	for {
		select {
		case <-tk.C:
			//tasks, err := m.GetScheduleTasks()
			tasks, err := m.GetAllTasks()
			if err != nil {
				log.Printf("scheduler get to schedule tasks err %v\n", err)
				continue
			}

			if len(tasks) > 0  {
				log.Printf("scheduler get %d to schedule tasks \n", len(tasks))
			}

			tasks = m.taskManager.ReduceReplicatedTask(tasks)
			if len(tasks) > 0 {
				log.Printf("scheduler[%s]: query prepare schedule tasks: %d\n", m.TaskType.String(), len(tasks))
			}

			m.Schedule(nil, tasks)
			if len(tasks) > 0 {
				log.Printf("<<scheduler manager(%s) scheduler %d task over>>", m.TaskType.String(), len(tasks))
			}
		}
	}
}
