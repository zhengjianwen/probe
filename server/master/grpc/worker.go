package grpc

import (
	"errors"
	"github.com/rongyungo/probe/server/master/types"
)

var ErrConnNotFound = errors.New("worker grpc stream not found")
var ErrConnExpired = errors.New("worker grpc stream expired")
var ErrConnClosed = errors.New("worker grpc stream closed")

// send a task to a worker channel
func (m *master) SendTask(wid int64, tk []types.TaskInterface) error {
	conn, err := m.getWorkerCon(wid)
	if err != nil {
		return err
	}

	if conn.isHealth() {
		// record channel may be full error
		go conn.recordMessage(tk...)
	}

	return nil
}

func (m *master) getWorkerCon(wid int64) (*conn, error) {
	con, ok := m.workerConnMap[wid]
	if !ok {
		return nil, ErrConnNotFound
	}

	return con, nil
}

func (m *master) reSend() {
	for _, conn := range m.workerConnMap {
		if conn.hasResendMessages() {
			//go conn.resendMessage()
		}
	}
}
