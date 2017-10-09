package worker

import (
	pb "github.com/rongyungo/probe/server/proto"
	"reflect"
)

func NewTaskManager() *taskManager {
	return &taskManager{
		m:             make(map[string]*pb.TaskInfo),
		execResultMap: make(map[string]*pb.TaskResult),
	}
}

type taskManager struct {
	m             map[string]*pb.TaskInfo
	execResultMap map[string]*pb.TaskResult
}

func (m *taskManager) AddTask(t *pb.TaskInfo) {
	if _, ok := m.m[t.TaskId]; !ok {
		m.m[t.TaskId] = t
	}
}

func change(a, b *pb.TaskInfo) bool {
	return reflect.DeepEqual(a, b)
}
