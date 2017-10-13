package apm

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func pushFtp(res *pb.TaskResult) error {
	return pushToApm(getDelayMetric(res), getCodeMetric(res))
}
