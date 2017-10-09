package exec

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func Execute(t *pb.TaskInfo) *pb.TaskResult {
	var res *pb.TaskResult
	switch t.Type {
	case pb.TaskInfo_HTTP:
		if t.Http_Spec == nil {
			return nil
		}
		res = ProbeHttp(t)
	case pb.TaskInfo_DNS:
		if t.Dns_Spec == nil {
			return nil
		}
		res = ProbeDns(t)
	case pb.TaskInfo_PING:
		if t.Ping_Spec == nil {
			return nil
		}
		res = ProbePing(t)
	case pb.TaskInfo_TRACE_ROUTE:
		if t.Traceroute_Spec == nil {
			return nil
		}
		res = ProbeTraceRoute(t)
	case pb.TaskInfo_TCP:
		if t.TcpSpec == nil {
			return nil
		}
		res = ProbeTcp(t)
	case pb.TaskInfo_UDP:
		if t.UdpSpec == nil {
			return nil
		}
		res = ProbeUdp(t)
	case pb.TaskInfo_FTP:
		if t.FtpSpec == nil {
			return nil
		}
		res = ProbeFtp(t)
	}

	res.ScheduleTime = t.ScheduleTime
	return res
}
