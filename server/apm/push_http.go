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

	mv1 := bufPool.Get().(model.MetricValue)
	defer bufPool.Put(mv1)

	mv1.Metric = getMetric(res.Type, "sc")
	mv1.Endpoint = fmt.Sprintf("url-%d-%d", wid, res.TaskId)
	mv1.Value = res.GetHttp().GetStatusCode()
	mv1.Timestamp = time.Now().Unix() - 60
	mv1.Type = "GAUGE"
	mv1.Step = int(res.GetPeriodSec())

	return pushToApm(&mv1, getWorkerDelayMetric(wid, res))
}

func PushHttpStat(tid int64, av, delay, step int) error {
	mav := getHttpTaskAvMetric(tid, av, step)
	mdelay := getHttpTaskDelayMetric(tid, delay, step)
	return pushToApm(mav, mdelay)
}

func getHttpTaskAvMetric(tid int64, av, step int) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)
	defer bufPool.Put(ret)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(pb.TaskType_HTTP, "av"), fmt.Sprintf("url-%d", tid), av
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", step

	return &ret
}

func getHttpTaskDelayMetric(tid int64, delay, step int) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)
	defer bufPool.Put(ret)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(pb.TaskType_HTTP, "delay"), fmt.Sprintf("url-%d", tid), delay
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", step

	return &ret
}
