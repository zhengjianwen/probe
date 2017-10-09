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

type conn struct {
	*sync.RWMutex

	workerId string
	status   string
	errorMsg string

	createTime      int64
	updateTime      int64
	healthCheckTime int64

	runSendMessage bool

	sendCh   chan *types.Task
	resendCh chan *pb.TaskInfo
	finalBuf []*pb.TaskInfo //5 minutes

	unHealthCKCh chan struct{}
	closeCh      chan struct{}
}

func newConn(wid string) *conn {
	now := time.Now().Unix()
	return &conn{
		workerId:        wid,
		healthCheckTime: now,
		createTime:      now,
		updateTime:      now,
		RWMutex:         new(sync.RWMutex),
		sendCh:          make(chan *types.Task, maxSendChanLength),
		resendCh:        make(chan *pb.TaskInfo, maxResendChanLength),
		finalBuf:        make([]*pb.TaskInfo, 1),
		unHealthCKCh:    make(chan struct{}, 5),
		closeCh:         make(chan struct{}, 2),
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

func (c *conn) addTaskBuf(t *pb.TaskInfo) {
	if t != nil {
		c.finalBuf = append(c.finalBuf, t)
	}
}

func (c *conn) emitTaskBuf() {
	for _, t := range c.finalBuf {
		c.resendCh <- t
	}
}

func (c *conn) Print() {
	log.Printf("worker(%s) conn  info:", c.workerId)
	log.Printf("health check: %s\n", time.Unix(c.healthCheckTime, 0).String())
	log.Printf("create time: %s ", time.Unix(c.createTime, 0).String())
	log.Printf("update time: %s", time.Unix(c.updateTime, 0).String())
}

func (c *conn) recordMessage(tk *types.Task) error {
	if len(c.sendCh) <= maxSendChanLength-1 {
		log.Printf("master conn propare sending task %s\n", tk.String())
		c.sendCh <- tk
		return nil
	}

	return ErrMessgeSendChFull
}

func (c *conn) reRecordMessage(t *pb.TaskInfo) {
	if len(c.resendCh) <= maxResendChanLength-10 {
		c.resendCh <- t
	} else {
		log.Printf("dropping a reRecord  message: %#v\n", *t)
	}
}

func (c conn) hasResendMessages() bool {
	return len(c.resendCh) > 0
}

// this is used to label a conn working
func (c *conn) updateTm() {
	c.updateTime = time.Now().Unix()
	c.healthCheckTime = c.updateTime
}

var unhealthCheck = pb.TaskInfo{Type: pb.TaskInfo_UNHEALTH_CHECK}

func (c *conn) sendMessage(stream pb.MasterWorker_SubscribeServer) {
	defer c.stopSendMessageFn()

	log.Printf("<<conn(%d, %d, %d) start sending message>>\n", len(c.sendCh), len(c.resendCh), len(c.finalBuf))
	for {
		select {
		case t := <-c.sendCh:
			topic := t.Convert()
			if err := stream.Send(topic); err != nil {
				c.setErrorMsg(err)
				c.setStatus(Status_Resending)
				c.reRecordMessage(topic)
			}
		case t := <-c.resendCh:
			if err := stream.Send(t); err != nil {
				c.setErrorMsg(err)
				c.setStatus(Status_Resend_Err)
				c.addTaskBuf(t)
			}
		case <-c.unHealthCKCh:
			if err := stream.Send(&unhealthCheck); err != nil {
				c.setErrorMsg(err)
				c.setStatus(Status_UnHealth_Err)
			}
		case <-c.closeCh:
			log.Println("<<conn stop sending message goroutine>>")
			return
		}
	}
}
