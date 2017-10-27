package apm

import (
	pb "github.com/rongyungo/probe/server/proto"
	"testing"
)

func TestPush(t *testing.T) {
	if err := PushWorker(1, &pb.TaskResult{
		TaskId:    900,
		Type:      pb.TaskType_HTTP,
		Http:      &pb.TaskResultHttp{400},
		PeriodSec: 60,
		DelayMs:   100,
		ErrorCode: 4,
	}); err != nil {
		t.Fatal(err)
	}
}
