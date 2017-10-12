package scheduler

import (
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"reflect"
)

func convertTasks(slice interface{}) []types.TaskInterface {
	var ret []types.TaskInterface

	sv := reflect.ValueOf(slice).Elem()
	for i := 0; i < sv.Len(); i++ {
		ret = append(ret, sv.Index(i).Interface().(types.TaskInterface))
	}

	return ret
}

func getSliceByType(tp pb.TaskType) interface{} {
	switch tp {
	case pb.TaskType_HTTP:
		var l []*types.Task_Http
		return &l
	case pb.TaskType_PING:
		var l []*types.Task_Ping
		return &l
	case pb.TaskType_TRACE_ROUTE:
		var l []*types.Task_TraceRoute
		return &l
	case pb.TaskType_TCP:
		var l []*types.Task_Tcp
		return &l
	case pb.TaskType_UDP:
		var l []*types.Task_Udp
		return &l
	case pb.TaskType_DNS:
		var l []*types.Task_Dns
		return &l
	case pb.TaskType_FTP:
		var l []*types.Task_Ftp
		return &l
	}
	return nil
}
