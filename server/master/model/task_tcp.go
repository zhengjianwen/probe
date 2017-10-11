package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Tcp(tk *types.Task_Tcp) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_Tcp(tid int64) (*types.Task_Tcp, error) {
	var task types.Task_Tcp
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var tt types.Task_Tcp

func DeleteTask_Tcp(id int64) error {
	_, err := Orm.Id(id).Delete(tt)
	return err
}
