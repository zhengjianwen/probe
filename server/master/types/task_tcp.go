package types

import (
	pb "github.com/rongyungo/probe/server/proto"
	"time"

	"fmt"
	errutil "github.com/rongyungo/probe/util/errors"
)

type Task_Tcp struct {
	pb.BasicInfo `xorm:"extends"`
	pb.TcpSpec   `xorm:"extends"`
}

func (t *Task_Tcp) String() string {
	var tabs string = "\t"

	return fmt.Sprintf("%d \t %s \t%s %d \t\t %s \t %s",
		t.Id, t.Type.String(), tabs, t.PeriodSec, time.Unix(t.CreateTime, 0).Format("2006-01-02 15:04:05"), t.TcpSpec.Host)

}
func (t *Task_Tcp) GetId() int64 {
	return t.Id
}

func (t *Task_Tcp) GetPeriodSec() int64 {
	return t.PeriodSec
}

func (t *Task_Tcp) Validate() error {
	if t.PeriodSec < MinPeriodSec {
		return errutil.ErrTaskPeriodTooLess
	}

	return nil
}

func (t *Task_Tcp) Complete() {
	now := time.Now().Unix()
	t.CreateTime, t.UpdateTime, t.ScheduleTime = now, now, now+int64(t.PeriodSec)
	t.Type = pb.TaskType_TCP
}

func (t *Task_Tcp) TableName() string {
	return "task_http"
}

func (t *Task_Tcp) Convert() *pb.Task {
	return &pb.Task{
		BasicInfo: &t.BasicInfo,
		TcpSpec:   &t.TcpSpec,
	}
}
