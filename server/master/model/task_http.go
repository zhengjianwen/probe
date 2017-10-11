package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Http(tk *types.Task_Http) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
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

var th types.Task_Http

func DeleteTask_Http(id int64) error {
	_, err := Orm.Id(id).Delete(th)
	return err
}
