package apm

import (
	"fmt"
	"github.com/rongyungo/apm/common/model"
	pb "github.com/rongyungo/probe/server/proto"
	"time"
)

func pushWorkerHttp(wid int64, res *pb.TaskResult) error {
	if res.GetHttp() == nil {
		return nil
	}

	tm := time.Now().Unix() - 60
	mv1, mv2 := bufPool.Get().(model.MetricValue), bufPool.Get().(model.MetricValue)
	defer bufPool.Put(mv1)
	defer bufPool.Put(mv2)

	mv1.Metric, mv2.Metric = getMetric(res.Type, "statuscode"), getMetric(res.Type, "error")
	mv1.Endpoint, mv2.Endpoint = fmt.Sprintf("url-%d-%d", wid, res.TaskId), fmt.Sprintf("url-%d-%d", wid, res.TaskId)
	mv1.Value, mv2.Value = res.GetHttp().GetStatusCode(), res.GetErrorCode()
	mv1.Timestamp, mv2.Timestamp = tm, tm
	mv1.Type, mv2.Type = "GAUGE", "GAUGE"
	mv1.Step, mv2.Step = int(res.GetPeriodSec()), int(res.GetPeriodSec())

	return pushToApm(&mv1, getWorkerDelayMetric(wid, res), &mv2)
}

func PushHttpStat(tid int64, av float64, step int) error {
	return pushToApm(getTaskHttpStatMetric(tid, av, step))
}

func getTaskHttpStatMetric(tid int64, av float64, step int) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(pb.TaskType_HTTP, "av"), fmt.Sprintf("url-%d", tid), av
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", step

	return &ret
}
