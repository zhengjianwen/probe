package types

import (
	pb "github.com/rongyungo/probe/server/proto"
	"time"

	"fmt"
	errutil "github.com/rongyungo/probe/util/errors"
)

type Task_Udp struct {
	pb.BasicInfo `xorm:"extends"`
	pb.UdpSpec   `xorm:"extends"`
}

func (t *Task_Udp) String() string {
	var tabs string = "\t"

	return fmt.Sprintf("%d \t %s \t%s %d \t\t %s \t %s",
		t.Id, t.Type.String(), tabs, t.PeriodSec, time.Unix(t.CreateTime, 0).Format("2006-01-02 15:04:05"), t.UdpSpec.Host)

}
func (t *Task_Udp) GetId() int64 {
	return t.Id
}

func (t *Task_Udp) GetPeriodSec() int64 {
	return t.PeriodSec
}

func (t *Task_Udp) Validate() error {
	if t.PeriodSec < MinPeriodSec {
		return errutil.ErrTaskPeriodTooLess
	}

	return nil
}

func (t *Task_Udp) Complete() {
	now := time.Now().Unix()
	t.CreateTime, t.UpdateTime, t.ScheduleTime = now, now, now+int64(t.PeriodSec)
	t.Type = pb.TaskType_UDP
}

func (t *Task_Udp) TableName() string {
	return "task_udp"
}

func (t *Task_Udp) Convert() *pb.Task {
	return &pb.Task{
		BasicInfo: &t.BasicInfo,
		UdpSpec:   &t.UdpSpec,
	}
}
