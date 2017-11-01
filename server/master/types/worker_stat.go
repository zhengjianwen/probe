package types

type TaskSchedule struct {
	TaskId      	int64 	`xorm:"task_id pk"`
	ScheduleTime 	int64 	`xorm:"schedule_time pk"`
	SuccessN 		int64
	ErrorN 			int64
	PeriodSec       int32
	IfStat          bool
}

func (p *TaskSchedule) TableName() string {
	return "task_schedule"
}
