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

type point struct {
	Metric    string `json:"metric"`
	Endpoint  string `json:"endpoint"`
	Tags      string `json:"tags"`
	Value     int    `json:"value"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	Step      int    `json:"step"`
}

func Push(res *pb.TaskResult) error {
	if res == nil {
		return nil
	}

	switch res.GetType() {
	case pb.TaskType_HTTP:
		if res.GetHttp() != nil {
			return pushHttp(res)
		}
	case pb.TaskType_DNS:
	case pb.TaskType_PING:
	case pb.TaskType_TRACE_ROUTE:
	case pb.TaskType_TCP:
	case pb.TaskType_UDP:
	case pb.TaskType_FTP:
	}
	return nil
}

func pushHttp(res *pb.TaskResult) error {
	tm := time.Now().Unix() - 60
	v1, v2, v3 := bufPool.Get().(model.MetricValue), bufPool.Get().(model.MetricValue), bufPool.Get().(model.MetricValue)
	defer bufPool.Put(v1)
	defer bufPool.Put(v2)
	defer bufPool.Put(v3)

	v1.Metric, v2.Metric, v3.Metric = getMetric(res.Type, "statuscode"), getMetric(res.Type, "delay"), getMetric(res.Type, "error")
	v1.Endpoint, v2.Endpoint, v3.Endpoint = fmt.Sprintf("url-%d", res.TaskId), fmt.Sprintf("url-%d", res.TaskId), fmt.Sprintf("url-%d", res.TaskId)
	v1.Value, v2.Value, v3.Value = res.GetHttp().GetStatusCode(), res.GetDelayMs(), res.GetErrorCode()
	v1.Timestamp, v2.Timestamp, v3.Timestamp = tm, tm, tm
	v1.Type, v2.Type, v3.Type = "GAUGE", "GAUGE", "GAUGE"
	v1.Step, v2.Step, v3.Step = int(res.GetPeriodSec()), int(res.GetPeriodSec()), int(res.GetPeriodSec())

	return push(&v1, &v2, &v3)
}

func push(vs ...*model.MetricValue) error {
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

func getMetric(tp pb.TaskType, key string) string {
	return fmt.Sprintf("url.%s.%s", strings.ToLower(tp.String()), key)
}
