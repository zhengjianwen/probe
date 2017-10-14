package model

import (
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/util/sql"
	"testing"
)

func Test_AdminEditWorker(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	wk := &types.Worker{Country: "中国", Province: "辽宁", City: "大连", Operator: "联通"}
	if err := AdminEditWorker(1, wk); err != nil {
		t.Fatal(err)
	}
}
