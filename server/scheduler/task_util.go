package scheduler

import (
	"container/list"
	"github.com/rongyungo/probe/server/master/types"
)

func loadTasksToList(arr []*types.Task, l *list.List) {
	for _, task := range arr {
		l.PushBack(task)
	}
}
