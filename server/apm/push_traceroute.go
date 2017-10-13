package apm

import (
	"fmt"
	"github.com/rongyungo/apm/common/model"
	pb "github.com/rongyungo/probe/server/proto"
	"time"
)

func pushTraceRoute(res *pb.TaskResult) error {
	if res.GetTraceroute() == nil {
		return nil
	}

	return pushToApm(getDelayMetric(res), getHopMetric(res), getErrHopMetric(res))
}

func getHopMetric(res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "hop"), fmt.Sprintf("url-%d", res.TaskId), res.Traceroute.Hops
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}

func getErrHopMetric(res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "errhop"), fmt.Sprintf("url-%d", res.TaskId), res.Traceroute.ErrHops
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}
