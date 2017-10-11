package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask_Ftp(tk *types.Task_Ftp) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask_Ftp(tid int64) (*types.Task_Ftp, error) {
	var task types.Task_Ftp
	if ok, err := Orm.Id(tid).Get(&task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskIdNotFound
	}
	return &task, nil
}

var tf types.Task_Ftp

func DeleteTask_Ftp(id int64) error {
	_, err := Orm.Id(id).Delete(tf)
	return err
}
