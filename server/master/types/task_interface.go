package types

import pb "github.com/rongyungo/probe/server/proto"

type TaskInterface interface {
	String() string
	GetId() int64
	GetPeriodSec() int64
	Convert() *pb.Task
	SetScheduleTime(st int64)
	AddRuleId(int64)
	RemoveRuleId(int64)
	GetType() pb.TaskType
}
