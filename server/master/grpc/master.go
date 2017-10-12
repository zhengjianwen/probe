package grpc

import (
	"github.com/rongyungo/probe/server/master/types"
	"sync"
	"time"

	pb "github.com/rongyungo/probe/server/proto"
	"log"
)

const (
	healthCheckSec         = 60
	tolerantHealthCheckSec = 5
)

var Master *master

type master struct {
	*sync.RWMutex

	// worker period sub second
	//subPeriodSec uint16

	// key:workerId, value:worker
	workersMap map[int64]*types.Worker

	// key:workerId, value: a writable gRpc stream to worker
	// when a new conn arrived, we should make sure the
	// worker's workerToTaskChMap task should send to new conn
	workerConnMap map[int64]*conn
}

func NewMaster() *master {
	m := master{
		RWMutex:       new(sync.RWMutex),
		workersMap:    make(map[int64]*types.Worker),
		workerConnMap: make(map[int64]*conn),
	}

	go m.initHouseKeeping()
	return &m
}

//quickly get all workers with no condition
func (m *master) GetWorkerIds() []int64 {
	m.RLock()
	defer m.RUnlock()

	var ids []int64
	for id := range m.workerConnMap {
		ids = append(ids, id)
	}

	return ids
}

func (m *master) initHouseKeeping() {
	for {
		select {
		case <-time.Tick(time.Second * 30):
			log.Printf("master start house keeping")
			m.cleanConn()
			log.Printf("master start house keeping over")
		}
	}
}

func (m *master) serveWorker(wId int64, ss pb.MasterWorker_SubscribeServer) chan<- struct{} {
	con := m.acceptConn(wId)
	//log.Printf("------------------------- %#v\n", con.finalBuf)
	// this part need reduce client multi connect

	if !con.ifSendMessageFnRun() {
		go con.sendMessage(ss)
	}

	return con.closeCh
}

// step1: worker connect to master
func (m *master) acceptConn(wId int64) *conn {
	m.Lock()
	defer m.Unlock()

	con, ok := m.workerConnMap[wId]
	if !ok {
		con = newConn(wId)
	}

	con.updateTm()
	con.open()

	m.workerConnMap[wId] = con
	return con
}

func (m *master) removeConn(wId int64) {
	m.Lock()
	defer m.Unlock()

	m.closeConnSession(wId)
	delete(m.workersMap, wId)
	delete(m.workerConnMap, wId)
}

func (m *master) closeConnSession(wId int64) {
	if con, ok := m.workerConnMap[wId]; ok {
		con.closeCh <- struct{}{}
	}
}

// @param wId := worker Id
// @param wt  := worker timestamp
func (m *master) healthCheck(wId int64) {
	m.RLock()
	defer m.RUnlock()

	con, ok := m.workerConnMap[wId]
	if ok {
		con.healthCheckTime = time.Now().Unix()
	}
}

func (m *master) cleanConn() {
	for workerId, conn := range m.workerConnMap {
		if m.isWorkerUnHealth(workerId) {
			log.Printf("master uninstall worker(%s) con\n", workerId)
			conn.Print()

			// TODO 主动监测
			if m.isWorkerDeath(workerId) {
				m.removeConn(workerId)
			}
		} else if len(conn.getStatus()) == 0 {
			conn.setStatus(Status_Health)
		}
		conn.Print()
	}
}

// no health report in a short time
func (m *master) isWorkerUnHealth(wId int64) bool {
	con, ok := m.workerConnMap[wId]
	if !ok {
		return true
	}

	return !con.isHealth()
}

// no health report in a long time, need lean memory
func (m *master) isWorkerDeath(wId int64) bool {
	con, ok := m.workerConnMap[wId]
	if !ok {
		return true
	}

	return time.Now().Unix()-time.Unix(con.healthCheckTime, 0).Unix() >= (5 * 60)
}
