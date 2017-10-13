package apm

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func pushDns(res *pb.TaskResult) error {
	return pushToApm(getDelayMetric(res))
}
