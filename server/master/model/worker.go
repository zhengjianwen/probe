package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func RegisterWorker(worker *types.Worker) error {
	_, err := Orm.Insert(worker)
	return err
}

//Country        string //国家
//Province       string //省
//City 		     string //市
//Operator       string //运营商
func AdminEditWorker(id int64, worker *types.Worker) error {
	if id <= 0 {
		return errutil.ErrWorkerIdEmpty
	}
	_, err := Orm.Id(id).Cols("country", "province", "city", "operator").Update(worker)
	return err
}

func AdminDelWorker(id int64) error {
	if id <= 0 {
		return errutil.ErrWorkerIdEmpty
	}

	_, err := Orm.Id(id).Delete(types.Worker{})
	return err
}

func ListWorkers(ids ...int64) (*[]types.Worker, error) {
	var l []types.Worker

	if len(ids) == 0 {
		return &l, Orm.Find(&l)
	} else {
		return &l, Orm.In("id", ids).Find(&l)
	}
}

func GetWorkerById(id int64) (*types.Worker, error) {
	var wk types.Worker
	if ok, err := Orm.Id(id).Get(&wk); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrWorkerNotFound
	}
	return &wk, nil
}