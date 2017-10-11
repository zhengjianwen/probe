package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Ping(tk *types.Task_Ping) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_Ping(tid int64) (*types.Task_Ping, error) {
	var task types.Task_Ping
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var tp types.Task_Ping

func DeleteTask_Ping(id int64) error {
	_, err := Orm.Id(id).Delete(tp)
	return err
}
