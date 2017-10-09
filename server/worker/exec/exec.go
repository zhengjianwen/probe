package exec

import (
	pb "github.com/ten-cloud/prober/server/proto"
	"log"
)

func Execute(t *pb.TaskInfo) *pb.TaskResult {
	switch t.Type {
	case pb.TaskInfo_HTTP:
		return ProbeHttp(t)
	case pb.TaskInfo_DNS:
		return ProbeDns(t)
	case pb.TaskInfo_PING:
		return ProbePing(t)
	case pb.TaskInfo_TRACE_ROUTE:
		return ProbeTraceRoute(t)
	default:
		log.Println("execute unknown task type %s \n", t.Type.String())
		return nil
	}
}
