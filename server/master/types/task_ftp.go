package types

import (
	pb "github.com/rongyungo/probe/server/proto"
	"time"

	"fmt"
	errutil "github.com/rongyungo/probe/util/errors"
)

type Task_Ftp struct {
	pb.BasicInfo `xorm:"extends"`
	pb.FtpSpec   `xorm:"extends"`
}

func (t *Task_Ftp) String() string {
	var tabs string = "\t"

	return fmt.Sprintf("%d \t %s \t%s %d \t\t %s \t %s",
		t.Id, t.Type.String(), tabs, t.PeriodSec, time.Unix(t.CreateTime, 0).Format("2006-01-02 15:04:05"), t.FtpSpec.Host)

}
func (t *Task_Ftp) GetId() int64 {
	return t.Id
}

func (t *Task_Ftp) GetPeriodSec() int64 {
	return t.PeriodSec
}

func (t *Task_Ftp) Validate() error {
	if t.PeriodSec < MinPeriodSec {
		return errutil.ErrTaskPeriodTooLess
	}

	return nil
}

func (t *Task_Ftp) Complete() {
	now := time.Now().Unix()
	t.CreateTime, t.UpdateTime, t.ScheduleTime = now, now, now+int64(t.PeriodSec)
	t.Type = pb.TaskType_FTP
}

func (t *Task_Ftp) TableName() string {
	return "task_ftp"
}

func (t *Task_Ftp) Convert() *pb.Task {
	return &pb.Task{
		BasicInfo: &t.BasicInfo,
		FtpSpec:   &t.FtpSpec,
	}
}
