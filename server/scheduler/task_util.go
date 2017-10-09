package scheduler

import (
	"container/list"
	"github.com/ten-cloud/prober/server/master/types"
)

func loadTasksToList(arr []*types.Task, l *list.List) {
	for _, task := range arr {
		l.PushBack(task)
	}
}
