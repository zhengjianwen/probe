package scheduler

import (
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
)

func (m *ScheduleManager) GetRunningTasks() ([]types.TaskInterface, error) {
	l := getSliceByType(m.TaskType)
	if err := m.Db.Where("stop = false AND type = ? ", m.TaskType).Find(l); err != nil {
		return nil, err
	}

	return convertTasks(l), nil
}

func (m *ScheduleManager) TableName() string {
	switch m.TaskType {
	case pb.TaskType_HTTP:
		return new(types.Task_Http).TableName()
	case pb.TaskType_DNS:
		return new(types.Task_Dns).TableName()
	case pb.TaskType_PING:
		return new(types.Task_Ping).TableName()
	case pb.TaskType_TRACE_ROUTE:
		return new(types.Task_TraceRoute).TableName()
	case pb.TaskType_TCP:
		return new(types.Task_Tcp).TableName()
	case pb.TaskType_UDP:
		return new(types.Task_Udp).TableName()
	case pb.TaskType_FTP:
		return new(types.Task_Ftp).TableName()
	}
	return "xxx"
}

func (m *ScheduleManager) CreateTaskSchedule(scheduleTime int64, workerN int, ts []types.TaskInterface) error {
	var ss []types.TaskSchedule
	for _, task := range ts {
		ss = append(ss, types.TaskSchedule{
			TaskType:     int64(task.GetType()),
			TaskId:       task.GetId(),
			ScheduleTime: scheduleTime,
			WorkerN:      workerN,
			PeriodSec:    task.GetPeriodSec(),
			OrgId:        task.GetOrgId(),
		})
	}

	_, err := m.Db.Insert(ss)
	return err
}
