package types

import "sort"

type TaskSchedule struct {
	TaskType     int64 `xorm:"task_type pk"`
	TaskId       int64 `xorm:"task_id pk"`
	ScheduleTime int64 `xorm:"schedule_time pk"`
	DelaySum     int64 `xorm:"delay_sum default 0"` //ms
	WorkerN      int
	SuccessN     int64 `xorm:"success_n default 0"`
	ErrorN       int64 `xorm:"error_n default 0"`
	PeriodSec    int64
	OrgId 		 int64 `xorm:"org_id"`
}

func (p *TaskSchedule) TableName() string {
	return "task_schedule"
}

func (p TaskSchedule) CanFinish() bool {
	return p.SuccessN+p.ErrorN == int64(p.WorkerN)
}

type TaskScheduleList []TaskSchedule

func (p TaskScheduleList) ReturnFinishedTask() *TaskSchedule {
	switch len(p) {
	case 0:
		return nil
	case 1:
		if p[0].CanFinish() {
			return &p[0]
		} else {
			return nil
		}
	case 2:
		sort.Sort(p)
		return &p[0]
	default:
		sort.Sort(p)
		if p[len(p)-1].CanFinish() {
			return &p[len(p)-1]
		} else {
			return &p[len(p)-2]
		}
	}
}

func (p TaskScheduleList) Len() int {
	return len(p)
}

func (p TaskScheduleList) Less(i, j int) bool {
	return p[i].ScheduleTime < p[j].ScheduleTime
}

func (p TaskScheduleList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type TaskStat struct {
	TaskType int64 `xorm:"task_type default 0"`
	TaskId   int64 `xorm:"task_id default 0"`
	OrgId 	 int64 `xorm:"org_id" default 0`
	DelaySum int64 `xorm:"delay_sum default 0"`
	SuccessN int64 `xorm:"success_n default 0"`
	ErrorN   int64 `xorm:"error_n default 0"`
}

func (p *TaskStat) TableName() string {
	return "task_stat"
}
