package types

import (
	"errors"
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
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
	TcpSpec        *pb.Task_Tcp        `json:"tcpSpec",bson:"tcpSpec"`
	UdpSpec        *pb.Task_Udp        `json:"udpSpec",bson:"udpSpec"`
	FtpSpec        *pb.Task_Ftp        `json:"ftpSpec",bson:"ftpSpec"`
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
		TcpSpec:         t.TcpSpec,
		UdpSpec:         t.UdpSpec,
		FtpSpec:         t.FtpSpec,
	}
}

func (t *Task) String() string {
	var url, tabs string = "", "\t"
	switch t.Type {
	case pb.TaskInfo_HTTP:
		url = t.HttpSpec.Url
	case pb.TaskInfo_DNS:
		url = t.DnsSpec.Domain
	case pb.TaskInfo_PING:
		url = t.PingSpec.Destination
	case pb.TaskInfo_TRACE_ROUTE:
		url, tabs = t.TraceRouteSpec.Destination, ""
	case pb.TaskInfo_TCP:
		url = fmt.Sprintf("%s:%d", t.TcpSpec.Host, t.TcpSpec.Port)
	case pb.TaskInfo_UDP:
		url = fmt.Sprintf("%s:%d", t.UdpSpec.Host, t.UdpSpec.Port)
	case pb.TaskInfo_FTP:
		url = fmt.Sprintf("%s:%d", t.FtpSpec.Host, t.FtpSpec.Port)
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

func (tk *Task) CreateComplete() {
	tk.Id = bson.NewObjectId()
	now := time.Now().Unix()

	tk.CreateTime, tk.UpdateTime, tk.ScheduleTime = now, now, now+int64(tk.PeriodSec)

	if tk.HttpSpec != nil {
		tk.Type = pb.TaskInfo_HTTP
	}
	if tk.DnsSpec != nil {
		tk.Type = pb.TaskInfo_DNS
	}
	if tk.PingSpec != nil {
		tk.Type = pb.TaskInfo_PING
		if tk.PingSpec.Timeout == 0 {
			tk.PingSpec.Timeout = 6
		}
	}
	if tk.TraceRouteSpec != nil {
		tk.Type = pb.TaskInfo_TRACE_ROUTE
	}

	if tk.TcpSpec != nil {
		tk.Type = pb.TaskInfo_TCP
	}

	if tk.UdpSpec != nil {
		tk.Type = pb.TaskInfo_UDP
	}

	if tk.FtpSpec != nil {
		tk.Type = pb.TaskInfo_FTP
	}
}
