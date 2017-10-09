package types

import (
	"errors"
	"fmt"
	pb "github.com/ten-cloud/prober/server/proto"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Task struct {
	Id             bson.ObjectId       `bson:"_id"`
	Type           pb.TaskInfoType     `json:"-",bson:"type"`
	HttpSpec       *pb.Task_Http       `json:"httpSpec",bson:"httpSpec"`
	DnsSpec        *pb.Task_Dns        `json:"dnsSpec",bson:"dnsSpec"`
	PingSpec       *pb.Task_Ping       `json:"pingSpec",bson:"pingSpec"`
	TraceRouteSpec *pb.Task_Traceroute `json:"traceRouteSpec",bson:"traceRouteSpec"`
	CreateTime     int64               `json:"-",bson:"createTime"`
	UpdateTime     int64               `json:"-",bson:"updateTime"`  //update task metadata, not a task operation
	ExecuteTime    int64               `json:"-",bson:"executeTime"` //a task executed time
	ScheduleTime   int64               `json:"-",bson:"scheduleTime"`
	PeriodSec      uint32              `json:"periodSec",bson:"periodSec"`
}

func (t Task) Convert() *pb.TaskInfo {
	return &pb.TaskInfo{
		TaskId:          t.Id.Hex(),
		Type:            t.Type,
		CreateTime:      t.CreateTime,
		UpdateTime:      t.UpdateTime,
		ExecuteTime:     t.ExecuteTime,
		ScheduleTime:    t.ScheduleTime,
		PeriodSec:       t.PeriodSec,
		Http_Spec:       t.HttpSpec,
		Dns_Spec:        t.DnsSpec,
		Ping_Spec:       t.PingSpec,
		Traceroute_Spec: t.TraceRouteSpec,
	}
}

func (t *Task) String() string {
	var url, tabs string
	switch t.Type {
	case pb.TaskInfo_HTTP:
		url, tabs = t.HttpSpec.Url, "\t"
	case pb.TaskInfo_DNS:
		url, tabs = t.DnsSpec.Domain, "\t"
	case pb.TaskInfo_PING:
		url, tabs = t.PingSpec.Destination, "\t"
	case pb.TaskInfo_TRACE_ROUTE:
		url = t.TraceRouteSpec.Destination
	}

	return fmt.Sprintf("%s \t %s \t%s %d \t\t %s \t %s",
		t.Id.Hex(), t.Type.String(), tabs, t.PeriodSec, time.Unix(t.CreateTime, 0).Format("2006-01-02 15:04:05"), url)
}

func (t *Task) Title() string {
	return fmt.Sprintf("TID \t\t\t\t TYPE \t\t PERIOD \t CT \t\t\t URL")
}

func (t *Task) Validate() error {
	if t.HttpSpec == nil {
		return errors.New("param httpSpec not found")
	}

	return Validate(t.HttpSpec)
}
