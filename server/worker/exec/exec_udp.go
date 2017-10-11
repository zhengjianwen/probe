package exec

import (
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
	"net"
	"strings"
	"time"
)

func ProbeUdp(t *pb.Task) *pb.TaskResult {
	var start int64 = time.Now().UnixNano()
	res := pb.TaskResult{
		TaskId:  t.BasicInfo.GetId(),
		Type:    t.BasicInfo.GetType(),
		StartMs: start / 1e6}

	con, err := net.Dial("udp", fmt.Sprintf("%s:%d", t.UdpSpec.Host, t.UdpSpec.Port))
	res.DelayMs = (time.Now().UnixNano() - start) / 1e6
	if err != nil {
		res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_NET_DIAL
		return &res
	}
	defer con.Close()

	if len(t.UdpSpec.ReqContent) > 0 {
		if _, err := con.Write(([]byte(t.UdpSpec.ReqContent))); err != nil {
			res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_UDP_REQUEST
			return &res
		}
	}
	if len(t.UdpSpec.ResMatchContent) > 0 {
		var data []byte
		if _, err := con.Read(data); err != nil {
			res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_UDP_RESPONSE
			return &res
		} else if len(t.UdpSpec.ResMatchContent) > 0 && !strings.Contains(string(data), t.UdpSpec.ResMatchContent) {
			res.Error, res.ErrorCode = "udp response unmatch", pb.TaskResult_ERR_UDP_RESPONSE_UNMATCH
			return &res
		}
	}

	res.Success = true
	return &res
}
