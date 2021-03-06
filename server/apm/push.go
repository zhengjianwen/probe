package apm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rongyungo/apm/common/model"
	pb "github.com/rongyungo/probe/server/proto"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var bufPool = sync.Pool{
	New: func() interface{} { return model.MetricValue{} },
}

var Conf struct {
	Url   string
	OrgId string
	Token string
}

func init() {
	Conf.Url = "http://www.opdeck.com"
	Conf.OrgId = "1"
	Conf.Token = "ui49hfowlx0wkxoe,cjeaiqoei93ms8mx821kx"
}

func PushWorker(wid int64, res *pb.TaskResult) error {
	if res == nil {
		return nil
	}

	switch res.GetType() {
	case pb.TaskType_HTTP:
		return pushWorkerHttp(wid, res)

	case pb.TaskType_DNS:
		return pushDns(res)

	case pb.TaskType_PING:
		return pushPing(res)

	case pb.TaskType_TRACE_ROUTE:
		return pushTraceRoute(res)

	case pb.TaskType_TCP:
		return pushTcp(res)

	case pb.TaskType_UDP:
		return pushUdp(res)

	case pb.TaskType_FTP:
		return pushFtp(res)
	}
	return nil
}

func pushToApm(vs ...*model.MetricValue) error {
	bs, err := json.Marshal(vs)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(bs)

	var resp *http.Response
	resp, err = http.Post(fmt.Sprintf("%s/v1/push?orgid=%s&token=%s", Conf.Url, Conf.OrgId, Conf.Token), "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.Body != nil {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			return errors.New(string(data))
		}
	}

	return nil
}

func pushToApmWithOrgId(oid int64, vs ...*model.MetricValue) error {
	bs, err := json.Marshal(vs)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(bs)

	url := fmt.Sprintf("%s/v1/push?orgid=%d&token=%s", Conf.Url, oid, Conf.Token)
	var resp *http.Response
	resp, err = http.Post(url, "application/json", buf)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if resp.Body != nil {
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)
			return errors.New(string(data))
		}
	}

	return nil
}

func getDelayMetric(res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "delay"), fmt.Sprintf("url-%d", res.TaskId), res.GetDelayMs()
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}

func getWorkerDelayMetric(wid int64, res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "delay"), fmt.Sprintf("url-%d-%d", wid, res.TaskId), res.GetDelayMs()
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}

func getCodeMetric(res *pb.TaskResult) *model.MetricValue {
	ret := bufPool.Get().(model.MetricValue)

	ret.Metric, ret.Endpoint, ret.Value = getMetric(res.Type, "code"), fmt.Sprintf("url-%d", res.TaskId), res.GetErrorCode()
	ret.Timestamp, ret.Type, ret.Step = time.Now().Unix()-60, "GAUGE", int(res.GetPeriodSec())

	return &ret
}

func getMetric(tp pb.TaskType, key string) string {
	return fmt.Sprintf("url.%s.%s", strings.ToLower(tp.String()), key)
}
