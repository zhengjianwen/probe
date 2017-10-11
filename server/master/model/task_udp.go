package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Udp(tk *types.Task_Udp) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_Udp(tid int64) (*types.Task_Udp, error) {
	var task types.Task_Udp
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var tu types.Task_Udp

func DeleteTask_Udp(id int64) error {
	_, err := Orm.Id(id).Delete(tu)
	return err
}
