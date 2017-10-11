package model

import (
	"github.com/rongyungo/probe/server/master/types"
)

func RegisterWorker(worker *types.Worker) error {
	_, err := Orm.Insert(worker)
	return err
}
