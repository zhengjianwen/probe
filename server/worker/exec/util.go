package exec

import (
	"github.com/miekg/dns"
	pb "github.com/ten-cloud/prober/server/proto"
	"net"
	"strings"
	"time"
)

func Return(tid string, err error, start int64) *pb.TaskResult {
	var res pb.TaskResult
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Success = true
	}

	res.DelayMs = (time.Now().UnixNano() - start) / 1e6
	res.TaskId = tid
	res.StartMs = start / 1e6

	return &res
}

func ReturnWithCode(tid string, err error, code pb.TaskResultCode, start int64) *pb.TaskResult {
	var res pb.TaskResult
	if err != nil {
		res.Error = err.Error()
	} else {
		res.Success = true
	}

	res.DelayMs = (time.Now().UnixNano() - start) / 1e6
	res.TaskId = tid
	res.ErrorCode = code
	res.StartMs = start / 1e6

	return &res
}

func countDelayMs(start int64) int64 {
	return (time.Now().UnixNano() - start) / 1e6
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
