package types

import (
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
)

var MinPeriodSec int64 = 60
var TaskTypeToStructMappings = map[pb.TaskType]interface{}{
	pb.TaskType_HTTP:        []Task_Http{},
	pb.TaskType_DNS:         []Task_Dns{},
	pb.TaskType_PING:        []Task_Ping{},
	pb.TaskType_TRACE_ROUTE: []Task_TraceRoute{},
	pb.TaskType_TCP:         []Task_Tcp{},
	pb.TaskType_UDP:         []Task_Udp{},
	pb.TaskType_FTP:         []Task_Ftp{},
}

type Task struct {
	Spec TaskInterface
}

func (t *Task) Title() string {
	return fmt.Sprintf("TID \t\t TYPE \t\t PERIOD \t CT \t\t\t URL")
}
