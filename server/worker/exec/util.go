package exec

import (
	"github.com/miekg/dns"
	pb "github.com/rongyungo/probe/server/proto"
	"net"
	"strings"
	"time"
)

func Return(tid int64, tp pb.TaskType, err error, start int64) *pb.TaskResult {
	return ReturnWithCode(tid, tp, err, start, 0)
}

func ReturnWithCode(tid int64, tp pb.TaskType, err error, start int64, code pb.TaskResultCode) *pb.TaskResult {
	res := pb.TaskResult{
		TaskId:    tid,
		Type:      tp,
		DelayMs:   (time.Now().UnixNano() - start) / 1e6,
		StartMs:   start / 1e6,
		ErrorCode: code,
	}
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Success = true
	}

	return &res
}

func isAnswerMatchStr(answer []dns.RR, str string) bool {
	for _, rr := range answer {
		if strings.Contains(rr.String(), str) {
			return true
		}
	}
	return false
}

func isStrArrMatchStr(a []string, s string) bool {
	for _, ele := range a {
		if strings.Contains(ele, s) {
			return true
		}
	}
	return false
}

func isNSArrMatchStr(a []*net.NS, s string) bool {
	for _, ele := range a {
		if strings.Contains(ele.Host, s) {
			return true
		}
	}

	return false
}
