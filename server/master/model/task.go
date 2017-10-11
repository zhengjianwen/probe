package model

import (
	errutil "github.com/rongyungo/probe/util/errors"
)

func CreateTask(tk interface{}) (int64, error) {
	return Orm.Insert(tk)
}

func GetTask(tp string, id int64) (interface{}, error) {
	switch tp {
	case "http":
		return GetTask_Http(id)
	case "dns":
		return GetTask_Dns(id)
	case "ping":
		return GetTask_Ping(id)
	case "trace_route":
		return GetTask_TraceRoute(id)
	case "tcp":
		return GetTask_Tcp(id)
	case "udp":
		return GetTask_Udp(id)
	case "ftp":
		return GetTask_Ftp(id)
	}
	return nil, errutil.ErrUnSupportTaskType
}

func DeleteTask(tp string, id int64) error {
	switch tp {
	case "http":
		return DeleteTask_Http(id)
	case "dns":
		return DeleteTask_Dns(id)
	case "ping":
		return DeleteTask_Ping(id)
	case "trace_route":
		return DeleteTask_TraceRoute(id)
	case "tcp":
		return DeleteTask_Tcp(id)
	case "udp":
		return DeleteTask_Udp(id)
	case "ftp":
		return DeleteTask_Ftp(id)
	}
	return errutil.ErrUnSupportTaskType
}
