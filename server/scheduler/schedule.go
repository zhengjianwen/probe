package scheduler

import (
	"github.com/rongyungo/probe/server/master/types"

	"github.com/rongyungo/probe/server/master/grpc"
	"log"
)

func (m *ScheduleManager) Schedule(wids []int64, s *types.Strategy, ts []types.TaskInterface) error {
	for _, wid := range wids {
		if err := grpc.Master.SendTask(wid, ts); err != nil {
			log.Printf("schedule manager send worker %s err %v\n", wid, err)
		}
	}
	return nil
}
