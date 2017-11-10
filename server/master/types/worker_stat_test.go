package types

import (
	"sort"
	"testing"
)

func TestTaskSchedule_Sort(t *testing.T) {
	l := []TaskSchedule{
		{TaskId: 1, ScheduleTime: 2000},
		{TaskId: 1, ScheduleTime: 100},
		{TaskId: 2, ScheduleTime: 1000},
	}

	sort.Sort(TaskScheduleList(l))
	if l[0].ScheduleTime != 100 {
		t.Fatal()
	}
}
