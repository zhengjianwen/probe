package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask(tk interface{}) (int64, error) {
	return Orm.Insert(tk)
}

func CreateTask_Http(tk *types.Task_Http) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask(tp string, id int64) (interface{}, error) {
	switch tp {
	case "http":
		return GetTask_Http(id)
	case "dns":
	case "ping":
	case "trace_route":
	case "tcp":
	case "udp":
	case "ftp":
	}
	return nil, errutil.ErrUnSupportTaskType
}

func GetTask_Http(tid int64) (*types.Task_Http, error) {
	var task types.Task_Http
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

func DeleteTask(tp string, id int64) error {
	switch tp {
	case "http":
		return DeleteTask_Http(id)
	case "dns":
	case "ping":
	case "trace_route":
	case "tcp":
	case "udp":
	case "ftp":
	}
	return errutil.ErrUnSupportTaskType
}

var th types.Task_Http

func DeleteTask_Http(id int64) error {
	_, err := Orm.Id(id).Delete(th)
	return err
}
