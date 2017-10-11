package scheduler

import (
	"fmt"
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"log"
	"time"
)

func (m *ScheduleManager) GetAllTasks() ([]types.TaskInterface, error) {
	var ret []types.TaskInterface
	switch m.TaskType {
	case pb.TaskType_HTTP:
		if tks, err := m.getAllHttpTasks(); err != nil {
			return nil, err
		} else {
			ret = convertHttpTasks(tks)
		}
	}

	return ret, nil
}

func (m *ScheduleManager) GetScheduleTasks() ([]types.TaskInterface, error) {
	var ret []types.TaskInterface
	switch m.TaskType {
	case pb.TaskType_HTTP:
		if tks, err := m.getScheduleHttpTasks(); err != nil {
			return nil, err
		} else {
			ret = convertHttpTasks(tks)
		}
	}

	return ret, nil
}

func (m *ScheduleManager) getScheduleHttpTasks() ([]types.Task_Http, error) {
	var res []types.Task_Http
	now := time.Now().Unix()
	if err := m.Db.Where("schedule_time < ? and ? <= schedule_time + period_sec", now, now).Find(&res); err != nil {
		return nil, err
	}

	return res, nil
}

func (m *ScheduleManager) getAllHttpTasks() ([]types.Task_Http, error) {
	var res []types.Task_Http
	if err := m.Db.Find(&res); err != nil {
		return nil, err
	}
	return res, nil
}

func convertHttpTasks(ht []types.Task_Http) []types.TaskInterface {
	var ret []types.TaskInterface

	for i := range ht {
		ret = append(ret, types.TaskInterface(&ht[i]))
	}

	return ret
}
func (m *ScheduleManager) CorrectScheduleTime() {
	now := time.Now().Unix()

	sql := fmt.Sprintf("UPDATE %s SET schedule_time = %d where schedule_time + %s.period_sec < %d", m.TableName(), now, m.TableName(), now)
	if rs, err := m.Db.Exec(sql); err != nil {
		log.Printf("scheduler manager(%s) correct schedule_time err %v\n", m.TaskType.String(), err)
	} else {
		n, err := rs.RowsAffected()
		log.Printf("scheduler manager(%s) correct %d schedule_time with err %v\n", m.TaskType.String(), n, err)
	}
}

func (m *ScheduleManager) TableName() string {
	switch m.TaskType {
	case pb.TaskType_HTTP:
		return new(types.Task_Http).TableName()
	case pb.TaskType_DNS:
		return new(types.Task_Http).TableName()
	case pb.TaskType_PING:
		return new(types.Task_Http).TableName()
	case pb.TaskType_TRACE_ROUTE:
		return new(types.Task_Http).TableName()
	case pb.TaskType_TCP:
		return new(types.Task_Http).TableName()
	case pb.TaskType_UDP:
		return new(types.Task_Http).TableName()
	case pb.TaskType_FTP:
		return new(types.Task_Http).TableName()
	}
	return "xxx"
}

func (m *ScheduleManager) UpdateTaskTime(tid int64) error {
	sql := fmt.Sprintf("UPDATE %s SET schedule_time = %d WHERE id = %d",
		m.TableName(), time.Now().Unix(), tid)
	_, err := m.Db.Exec(sql)
	return err
}
