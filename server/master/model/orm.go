package model

import (
	"github.com/go-xorm/xorm"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/util/sql"
)

var Orm *xorm.Engine

func InitMySQL(cfg *sql.DatabaseConfig) (err error) {
	Orm, err = sql.InitMySQL(cfg)
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
	)
}
