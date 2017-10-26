package model

import (
	"github.com/go-xorm/xorm"
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/util/sql"
)

var Orm *xorm.Engine

func InitMySQL(cfg *sql.DatabaseConfig) (err error) {
	Orm, err = sql.InitMySQL(cfg)
	if err != nil {
		return err
	}
	return Sync()
}

func Sync() error {
	return Orm.Sync2(
		new(types.Worker),
		new(types.Task_Ftp),
		new(types.Task_Tcp),
		new(types.Task_Udp),
		new(types.Task_Dns),
		new(types.Task_Http),
		new(types.Task_Ping),
		new(types.Task_TraceRoute),
		new(pb.TaskResult),
		new(types.TaskSchedule),
	)
}
