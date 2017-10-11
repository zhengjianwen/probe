package exec

import (
	lib "github.com/kanocz/tracelib"
	pb "github.com/rongyungo/probe/server/proto"
	"time"
)

var cache *lib.LookupCache = lib.NewLookupCache()

func ProbeTraceRoute(t *pb.Task) *pb.TaskResult {
	start := time.Now().UnixNano()
	res := pb.TaskResult{
		TaskId:  t.BasicInfo.GetId(),
		Type:    t.BasicInfo.GetType(),
		StartMs: start / 1e6,
		DelayMs: (time.Now().UnixNano() - start) / 1e6,
	}

	hops, err := lib.RunTrace(t.TracerouteSpec.Destination, "0.0.0.0", time.Second, 64, cache, nil)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Success = true
	}

	res.Traceroute = &pb.TaskResultTraceroute{
		Hops:    int32(len(hops)),
		ErrHops: int32(len(getHopsErr(hops))),
	}

	return &res
}

func getHopsErr(h []lib.Hop) []error {
	var es []error
	for _, hp := range h {
		if hp.Error != nil {
			es = append(es, hp.Error)
		}
	}
	return es
}
