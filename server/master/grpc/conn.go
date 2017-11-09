package grpc

import (
	"sync"
	"time"

	"errors"
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"log"
)

const (
	maxSendChanLength   = 200
	maxResendChanLength = 400
)

var ErrMessgeSendChFull = errors.New("conn send channel fulled")
var Mod string = "dev"

var (
	Status_Health       string = "health"
	Status_UnHealth     string = "unhealth"
	Status_Resending    string = "resend"
	Status_Resend_Err   string = "resend error"
	Status_UnHealth_Err string = "unhealth error"
)

// a health conn means server accept health check reporter
// among the period of health, conn may lose it by a error
// scheduler can send task to a unhealth conn
// housekeeping will remove unhealth conn

type conn struct {
	*sync.RWMutex

	workerId int64
	status   string
	errorMsg string

	createTime      int64
	updateTime      int64
	healthCheckTime int64

	runSendMessage bool

	sendCh   chan types.TaskInterface
	resendCh chan types.TaskInterface
	//finalBuf []*pb.TaskInfo //5 minutes

	unHealthCKCh chan struct{}
	closed       bool //if closed, no data will send to sendCh
	closeCh      chan struct{}
}

func newConn(wid int64) *conn {
	now := time.Now().Unix()
	return &conn{
		workerId:        wid,
		healthCheckTime: now,
		createTime:      now,
		updateTime:      now,
		RWMutex:         new(sync.RWMutex),
		sendCh:          make(chan types.TaskInterface, maxSendChanLength),
		//resendCh:        make(chan *pb.TaskInfo, maxResendChanLength),
		//finalBuf:        make([]*pb.TaskInfo, 1),
		unHealthCKCh: make(chan struct{}, 5),
		closed:       false,
		closeCh:      make(chan struct{}, 2),
	}
}

func (c *conn) setStatus(s string) {
	c.status = s
}

func (c *conn) getStatus() string {
	return c.status
}

func (c *conn) setErrorMsg(e error) {
	c.errorMsg = e.Error()
}

func (c *conn) stopSendMessageFn() {
	c.Lock()
	defer c.Unlock()
	c.runSendMessage = false
}

func (c *conn) ifSendMessageFnRun() bool {
	c.RLock()
	defer c.RUnlock()
	return c.runSendMessage
}

func (c *conn) addTaskBuf(t *pb.Task) {
	if t != nil {
		//c.finalBuf = append(c.finalBuf, t)
	}
}

func (c *conn) close() {
	c.Lock()
	defer c.Unlock()

	c.closed = true
}

func (c *conn) open() {
	c.Lock()
	defer c.Unlock()

	c.closed = false
}

func (c *conn) isHealth() bool {
	// no health report in a short time
	return time.Now().Unix()-time.Unix(c.healthCheckTime, 0).Unix() < (healthCheckSec*2 + tolerantHealthCheckSec)
}

func (c *conn) emitTaskBuf() {
	//for _, t := range c.finalBuf {
	//	c.resendCh <- t
	//}
}

func (c *conn) Print() {
	log.Printf("worker(%d) conn  info:", c.workerId)
	log.Printf("health check: %s\n", time.Unix(c.healthCheckTime, 0).String())
	log.Printf("create time: %s ", time.Unix(c.createTime, 0).String())
	log.Printf("update time: %s", time.Unix(c.updateTime, 0).String())
}

func (c *conn) recordMessage(scheduleTime int64, tks ...types.TaskInterface) {
	for _, tk := range tks {
		if len(c.sendCh) >= maxSendChanLength-1 {
			log.Println("master grpc conn send channel full")
		}
		tk.SetScheduleTime(scheduleTime)
		c.sendCh <- tk
	}
}

func (c *conn) reRecordMessage(t *pb.Task) {
	//if len(c.resendCh) <= maxResendChanLength-10 {
	//	c.resendCh <- t
	//} else {
	//	log.Printf("dropping a reRecord  message: %#v\n", *t)
	//}
}

func (c conn) hasResendMessages() bool {
	//return len(c.resendCh) > 0
	return true
}

// this is used to label a conn working
func (c *conn) updateTm() {
	c.updateTime = time.Now().Unix()
	c.healthCheckTime = c.updateTime
}

//var unhealthCheck = pb.TaskInfo{Type: pb.TaskInfo_UNHEALTH_CHECK}

func (c *conn) sendMessage(stream pb.MasterWorker_SubscribeServer) {
	defer c.stopSendMessageFn()

	//log.Printf("<<conn(%d, %d, %d) start sending message>>\n", len(c.sendCh), len(c.resendCh), len(c.finalBuf))
	for {
		select {
		case t := <-c.sendCh:
			topic := t.Convert()
			if err := stream.Send(topic); err != nil {
				c.setErrorMsg(err)
				c.setStatus(Status_Resending)
				c.reRecordMessage(topic)
			}
		//case t := <-c.resendCh:
		//	if err := stream.Send(t); err != nil {
		//		c.setErrorMsg(err)
		//		c.setStatus(Status_Resend_Err)
		//		c.addTaskBuf(t)
		//	}
		//case <-c.unHealthCKCh:
		//	if err := stream.Send(&unhealthCheck); err != nil {
		//		c.setErrorMsg(err)
		//		c.setStatus(Status_UnHealth_Err)
		//	}
		case <-c.closeCh:
			c.close()
			log.Println("<<conn stop sending message goroutine>>")
			return
		}
	}
}
