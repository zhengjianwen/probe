package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_TraceRoute(tk *types.Task_TraceRoute) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_TraceRoute(tid int64) (*types.Task_TraceRoute, error) {
	var task types.Task_TraceRoute
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var ttr types.Task_TraceRoute

func DeleteTask_TraceRoute(id int64) error {
	_, err := Orm.Id(id).Delete(ttr)
	return err
}
