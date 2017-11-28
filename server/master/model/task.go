package model

import (
	"fmt"
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask(tk interface{Complete()}) (int64, error) {
	tk.Complete()
	return Orm.Insert(tk)
}

func GetTask(orgId, tid int64, tp string) (interface{}, error) {
	task := NewTaskPtr(tp)
	if ok, err := Orm.Where("id = ? AND org_id = ?", tid, orgId).Get(task); err != nil {
		return nil, err
	} else if !ok {
		return nil, errutil.ErrTaskNotFound
	}

	return task, nil
}

func GetOrgTask(orgId, nodeId int64, tp string) (interface{}, error) {
	l := NewTaskListPtr(tp)

	if nodeId == 0 {
		if err := Orm.Where("org_id = ?", orgId).Find(l); err != nil {
			return nil, err
		}
	} else {
		if err := Orm.Where("org_id = ? AND node_id = ?", orgId, nodeId).Find(l); err != nil {
			return nil, err
		}
	}

	return l, nil
}

func UpdateTask(orgId, tid int64, task interface{}) error {
	_, err := Orm.Where("id = ? AND org_id = ?", tid, orgId).Omit().Update(task)
	return err
}

func UpdateTaskRuleId1(orgId, tid int64, task interface{}) error {
	_, err := Orm.Where("id = ? AND org_id = ?", tid, orgId).Update(task)
	return err
}

func UpdateTaskRuleId2(orgId, tid int64, tp string, task interface {
	GetRuleIds() []int64
}) error {
	if len(task.GetRuleIds()) == 0 {
		sql := fmt.Sprintf("UPDATE %s SET rule_ids = NULL where id = %d AND org_id = %d",
			fmt.Sprintf("task_%s", tp), tid, orgId)
		_, err := Orm.Exec(sql)
		return err
	} else {
		return UpdateTaskRuleId1(orgId, tid, task)
	}
}

func DeleteTask(orgId, tid int64, tp string) error {
	task := NewTaskPtr(tp)
	_, err := Orm.Where("id = ? AND org_id = ?", tid, orgId).Delete(task)
	return err
}

func NewTaskPtr(tp string) interface{} {
	switch tp {
	case "http":
		return &types.Task_Http{}
	case "dns":
		return &types.Task_Dns{}
	case "ping":
		return &types.Task_Ping{}
	case "trace_route":
		return &types.Task_TraceRoute{}
	case "tcp":
		return &types.Task_Tcp{}
	case "udp":
		return &types.Task_Udp{}
	case "ftp":
		return &types.Task_Ftp{}
	}

	return nil
}

func NewTaskListPtr(tp string) interface{} {
	switch tp {
	case "http":
		return &[]types.Task_Http{}
	case "dns":
		return &[]types.Task_Dns{}
	case "ping":
		return &[]types.Task_Ping{}
	case "trace_route":
		return &[]types.Task_TraceRoute{}
	case "tcp":
		return &[]types.Task_Tcp{}
	case "udp":
		return &[]types.Task_Udp{}
	case "ftp":
		return &[]types.Task_Ftp{}
	}

	return nil
}
