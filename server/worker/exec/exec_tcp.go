package exec

import (
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
	"net"
	"time"
)

func ProbeTcp(t *pb.Task) *pb.TaskResult {
	start := time.Now().UnixNano()
	res := pb.TaskResult{
		TaskId:  t.BasicInfo.GetId(),
		Type:    t.BasicInfo.GetType(),
		StartMs: start / 1e6}

	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", t.TcpSpec.Host, t.TcpSpec.Port))
	if res.DelayMs = (time.Now().UnixNano() - start) / 1e6; err != nil {
		res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_NET_DIAL
	} else {
		defer con.Close()
		res.Success, res.ErrorCode = true, pb.TaskResult_OK
	}

	return &res
}
