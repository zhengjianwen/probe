package model

import (
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
	"time"
)

func RegisterWorker(worker *types.Worker) error {
	var wk types.Worker
	exist, err := Orm.Table("worker").Where("id = ?", worker.Id).Get(&wk)
	if err != nil {
		return err
	}

	if exist {
		if wk.Password != worker.Password {
			err = errutil.ErrWorkerUnAuthorized
		}
	} else {
		_, err = Orm.InsertOne(worker)
	}
	return err
}

func UpdateWorkerTime(ids ...int64) error {
	if len(ids) > 0 {
		_, err := Orm.In("id", ids).Update(types.Worker{
			UpdateTimestamp: 	time.Now().Unix(),
			Status: 			"ok",
		})
		return err
	}
	return nil
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
	return &l, Orm.In("id", ids).Where("status = 'ok' AND (UNIX_TIMESTAMP() - update_timestamp) < 300").Find(&l)
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