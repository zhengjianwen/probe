package scheduler

import (
	"container/list"
	"errors"
	"github.com/rongyungo/probe/server/master/types"
	"log"
	"os"
	"sync"
	"time"
)

type taskManager struct {
	*sync.RWMutex

	//map is used for fast delete
	//worker has no idea of when should execute a task
	//they are only the executor.
	taskMap  map[int64]*list.Element
	taskList *list.List

	updateTaskCh chan *types.TaskInterface

	scheduleMap map[int64]*taskRecord
}

func NewTaskManager() *taskManager {
	return &taskManager{
		RWMutex:      &sync.RWMutex{},
		taskMap:      make(map[int64]*list.Element),
		taskList:     list.New(),
		updateTaskCh: make(chan *types.TaskInterface, 200),
		scheduleMap:  make(map[int64]*taskRecord),
	}
}

func (m *taskManager) SyncTask(l []types.TaskInterface) error {
	m.Lock()
	defer m.Unlock()

	for _, taskI := range l {
		ele := m.taskList.PushBack(taskI)
		m.taskMap[taskI.GetId()] = ele
	}

	return nil
}

//bad lock to add task
func (m *taskManager) AddTask(tk types.TaskInterface) error {
	m.Lock()
	defer m.Unlock()

	if _, exist := m.taskMap[tk.GetId()]; exist {
		return errors.New("task already exists")
	}

	ele := m.taskList.PushBack(tk)
	m.taskMap[tk.GetId()] = ele

	return nil
}

func (m *taskManager) DelTask(tid int64) {
	m.Lock()
	defer m.Unlock()

	ele, ok := m.taskMap[tid]
	if ok {
		delete(m.taskMap, tid)
		m.taskList.Remove(ele)
	}
}

type taskRecord struct {
	task types.TaskInterface
	end  int64
}

func (t taskRecord) expired() bool {
	return t.end <= time.Now().Unix()
}

func (m *taskManager) ReduceReplicatedTask(l []types.TaskInterface) []types.TaskInterface {
	now := time.Now().Unix()
	var ret []types.TaskInterface
	for _, tk := range l {
		if history, ok := m.scheduleMap[tk.GetId()]; !ok {
			m.scheduleMap[tk.GetId()] = &taskRecord{
				task: tk,
				end:  now + tk.GetPeriodSec(),
			}
			ret = append(ret, tk)
		} else {
			if history.expired() {
				history.end = now + tk.GetPeriodSec()
				ret = append(ret, tk)
			}
		}
	}

	return ret
}

//更新任务涉及到较多组件， 目前先不支持
func (m *taskManager) UpdateTask(tk *types.TaskInterface) {
	m.updateTaskCh <- tk
}

func (m *taskManager) GetTask(tid int64) *types.Task {
	m.RLock()
	defer m.RUnlock()

	ele, ok := m.taskMap[tid]
	if ok {
		return ele.Value.(*types.Task)
	}

	return nil
}

func (m *taskManager) GetTaskIds() []int64 {
	m.RLock()
	defer m.RUnlock()

	var ids []int64
	for id := range m.taskMap {
		ids = append(ids, id)
	}
	return ids
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)
}

func (m *taskManager) StatTasks() {
	logger.Printf("stat task manager: total %d\n", m.taskList.Len())
	t := new(types.Task)

	logger.Println(t.Title())

	for e := m.taskList.Front(); e != nil; e = e.Next() {
		logger.Println(e.Value.(types.TaskInterface).String())
	}
}
