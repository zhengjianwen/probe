package types

import (
	pb "github.com/rongyungo/probe/server/proto"
	"time"

	"fmt"
	errutil "github.com/rongyungo/probe/util/errors"
)

type Task_Dns struct {
	pb.BasicInfo `xorm:"extends"`
	pb.DnsSpec   `xorm:"extends"`
}

func (t *Task_Dns) String() string {
	var tabs string = "\t"

	return fmt.Sprintf("%d \t %s \t%s %d \t\t %s \t %s",
		t.Id, t.BasicInfo.Type.String(), tabs, t.PeriodSec, time.Unix(t.CreateTime, 0).Format("2006-01-02 15:04:05"), t.DnsSpec.Domain)
}

func (t *Task_Dns) GetId() int64 {
	return t.Id
}

func (t *Task_Dns) GetPeriodSec() int64 {
	return t.PeriodSec
}

func (t *Task_Dns) Validate() error {
	if t.PeriodSec < MinPeriodSec {
		return errutil.ErrTaskPeriodTooLess
	}

	return nil
}

func (t *Task_Dns) Complete() {
	now := time.Now().Unix()
	t.CreateTime, t.UpdateTime, t.ScheduleTime = now, now, now+int64(t.PeriodSec)
	t.BasicInfo.Type = pb.TaskType_DNS
}

func (t *Task_Dns) TableName() string {
	return "task_dns"
}

func (t *Task_Dns) Convert() *pb.Task {
	return &pb.Task{
		BasicInfo: &t.BasicInfo,
		DnsSpec:   &t.DnsSpec,
	}
}
