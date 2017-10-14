package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask(tk interface{}) (int64, error) {
	v, _ := tk.(interface {
		Complete()
	})

	v.Complete()

	return Orm.Insert(tk)
}

func GetTask(tp string, id int64) (interface{}, error) {
	task := GetTypeStructPtr(tp)
	if ok, err := Orm.Id(id).Get(task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskNotFound
	}

	return task, nil
}

func UpdateTask(id int64, task interface{}) error {
	_, err := Orm.Id(id).Update(task)
	return err
}

func DeleteTask(tp string, id int64) error {
	task := GetTypeStructPtr(tp)
	_, err := Orm.Id(id).Delete(task)
	return err
}

func GetTypeStructPtr(tp string) interface{} {
	switch tp {
	case "http":
		return &types.Task_Http{}
	case "dns":
		return &types.Task_Dns{}
	case "ping":
		return &types.Task_Ping{}
	case "trace_route":
		return &types.Task_TraceRoute{}
	case "tcp":
		return &types.Task_Tcp{}
	case "udp":
		return &types.Task_Udp{}
	case "ftp":
		return &types.Task_Ftp{}
	}

	return nil
}
