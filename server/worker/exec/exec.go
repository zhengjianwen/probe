package exec

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func Execute(t *pb.Task) *pb.TaskResult {
	var res *pb.TaskResult
	if t.GetBasicInfo().PeriodSec <= 10 {
		return nil
	}

	switch t.BasicInfo.GetType() {
	case pb.TaskType_HTTP:
		if t.HttpSpec == nil {
			return nil
		}
		res = ProbeHttp(t)
	case pb.TaskType_DNS:
		if t.DnsSpec == nil {
			return nil
		}
		res = ProbeDns(t)
	case pb.TaskType_PING:
		if t.PingSpec == nil {
			return nil
		}
		res = ProbePing(t)
	case pb.TaskType_TRACE_ROUTE:
		if t.TracerouteSpec == nil {
			return nil
		}
		res = ProbeTraceRoute(t)
	case pb.TaskType_TCP:
		if t.TcpSpec == nil {
			return nil
		}
		res = ProbeTcp(t)
	case pb.TaskType_UDP:
		if t.UdpSpec == nil {
			return nil
		}
		res = ProbeUdp(t)
	case pb.TaskType_FTP:
		if t.FtpSpec == nil {
			return nil
		}
		res = ProbeFtp(t)
	}


	res.PeriodSec = t.GetBasicInfo().PeriodSec
	res.ScheduleTime = t.BasicInfo.GetScheduleTime()

	return res
}
