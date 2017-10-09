package scheduler

import (
	"container/list"
	"errors"
	"github.com/ten-cloud/prober/server/master/model"
	"github.com/ten-cloud/prober/server/master/types"
	"log"
	"os"
	"sync"
	"time"
)

var m *taskManager

type taskManager struct {
	*sync.RWMutex

	//map is used for fast delete
	//worker has no idea of when should execute a task
	//they are only the executor.
	taskMap  map[string]*list.Element
	taskList *list.List

	updateTaskCh chan *types.Task

	scheduleMap map[string]*taskHistory
}

func InitTaskManager() error {
	m = &taskManager{
		RWMutex:      &sync.RWMutex{},
		taskMap:      make(map[string]*list.Element),
		taskList:     list.New(),
		updateTaskCh: make(chan *types.Task, 200),
		scheduleMap:  make(map[string]*taskHistory),
	}

	return SyncTask()
}

func SyncTask() error {
	l, err := model.GetAllTasks()
	if err != nil {
		return err
	}

	m.Lock()
	defer m.Unlock()

	for _, task := range l {
		ele := m.taskList.PushBack(task)
		m.taskMap[task.Id.Hex()] = ele
	}

	return nil
}

//bad lock to add task
func AddTask(tk *types.Task) error {
	m.Lock()
	defer m.Unlock()

	if _, exist := m.taskMap[tk.Id.Hex()]; exist {
		return errors.New("task already exists")
	}

	ele := m.taskList.PushBack(tk)
	m.taskMap[tk.Id.Hex()] = ele

	return nil
}

func DelTask(tid string) {
	m.Lock()
	defer m.Unlock()

	ele, ok := m.taskMap[tid]
	if ok {
		delete(m.taskMap, tid)
		m.taskList.Remove(ele)
	}
}

type taskHistory struct {
	task *types.Task
	end  int64
}

func (t taskHistory) expired() bool {
	return t.end <= time.Now().Unix()
}

func ReduceReplicatedTask(l []*types.Task) []*types.Task {
	now := time.Now().Unix()
	var ret []*types.Task
	for _, tk := range l {
		if history, ok := m.scheduleMap[tk.Id.Hex()]; !ok {
			m.scheduleMap[tk.Id.Hex()] = &taskHistory{
				task: tk,
				end:  now + int64(tk.PeriodSec),
			}
			ret = append(ret, tk)
		} else {
			if history.expired() {
				history.end = now + int64(tk.PeriodSec)
				ret = append(ret, tk)
			}
		}
	}

	return ret
}

//更新任务涉及到较多组件， 目前先不支持
func UpdateTask(tk *types.Task) {
	m.updateTaskCh <- tk
}

func GetTask(tid string) *types.Task {
	m.RLock()
	defer m.RUnlock()

	ele, ok := m.taskMap[tid]
	if ok {
		return ele.Value.(*types.Task)
	}

	return nil
}

func GetTaskIds() []string {
	m.RLock()
	defer m.RUnlock()

	var ids []string
	for id := range m.taskMap {
		ids = append(ids, id)
	}
	return ids
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func StatTasks() {
	logger.Printf("stat task manager: total %d\n", m.taskList.Len())
	t := new(types.Task)
	logger.Println(t.Title())

	for e := m.taskList.Front(); e != nil; e = e.Next() {
		logger.Println(e.Value.(*types.Task).String())
	}
}
