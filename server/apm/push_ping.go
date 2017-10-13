package apm

import (
	"fmt"
	"github.com/rongyungo/apm/common/model"
	pb "github.com/rongyungo/probe/server/proto"
	"time"
)

func pushPing(res *pb.TaskResult) error {
	if res.GetPing() == nil {
		return nil
	}
	return pushToApm(getDelayMetric(res), getPingLostPercentMetric(res))
}

func getPingLostPercentMetric(res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "lost"), fmt.Sprintf("url-%d", res.TaskId), res.GetPing().GetLost()
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}
