package scheduler

import (
	"github.com/rongyungo/probe/server/master/grpc"
	"github.com/rongyungo/probe/server/master/types"
	"log"
	"time"
)

func (m *ScheduleManager) Schedule(s *types.Strategy, ts []types.TaskInterface) error {
	scheduleId := time.Now().Unix()

	wIds := grpc.Master.GetWorkerIds()
	if len(wIds) == 0 {
		return nil
	}

	if err := m.CreateTaskSchedule(scheduleId, len(wIds), ts); err != nil {
		return err
	}

	for _, wid := range wIds {
		if err := grpc.Master.SendTask(wid, ts, scheduleId); err != nil {
			log.Printf("schedule manager send worker %s err %v\n", wid, err)
		}
	}

	return nil
}
