package exec

import (
	lib "github.com/kanocz/tracelib"
	pb "github.com/ten-cloud/prober/server/proto"
	"time"
)

var cache *lib.LookupCache = lib.NewLookupCache()

func ProbeTraceRoute(t *pb.TaskInfo) *pb.TaskResult {
	if t.Traceroute_Spec == nil {
		return nil
	}

	var res pb.TaskResult
	start := time.Now().UnixNano()
	hops, err := lib.RunTrace(t.Traceroute_Spec.Destination, "0.0.0.0", time.Second, 64, cache, nil)
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Success = true
	}

	res.TaskId = t.TaskId
	res.StartMs = start / 1e6
	res.DelayMs = (time.Now().UnixNano() - start) / 1e6
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
