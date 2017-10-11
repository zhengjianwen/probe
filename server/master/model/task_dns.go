package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Dns(tk *types.Task_Dns) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_Dns(tid int64) (*types.Task_Dns, error) {
	var task types.Task_Dns
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var td types.Task_Dns

func DeleteTask_Dns(id int64) error {
	_, err := Orm.Id(id).Delete(td)
	return err
}
