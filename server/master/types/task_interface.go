package types

import pb "github.com/rongyungo/probe/server/proto"

type TaskInterface interface {
	GetId() int64
	GetPeriodSec() int64
	GetUrl() string
	GetType() pb.TaskType
	GetOrgId() int64

	String() string
	Convert() *pb.Task

	SetScheduleTime(st int64)
	AddRuleId(int64)
	RemoveRuleId(int64)
}
