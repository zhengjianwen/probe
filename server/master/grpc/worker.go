package grpc

import (
	"errors"
	"github.com/ten-cloud/prober/server/master/types"
)

var ErrConnNotFound = errors.New("worker grpc stream not found")
var ErrConnExpired = errors.New("worker grpc stream expired")
var ErrConnClosed = errors.New("worker grpc stream closed")

// send a task to a worker channel
func (m *master) SendTask(wid string, tk *types.Task) error {
	c, err := m.getWorkerCon(wid)
	if err != nil {
		return err
	}
	if err := c.recordMessage(tk); err != nil {
		return err
	}
	return nil
}

func (m *master) getWorkerCon(wid string) (*conn, error) {
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
