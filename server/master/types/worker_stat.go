package types

type TaskSchedule struct {
	TaskType 		int64   `xorm:"task_type pk"`
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

type TaskStat struct {
	TaskType    int64
	TaskId   	int64
	SuccessN  	int64
	ErrorN 		int64
}

func (p *TaskStat) TableName() string {
	return "task_stat"
}