package model

import (
	"testing"
	"github.com/rongyungo/probe/util/sql"
	"time"
)

func TestSyncWorkerStat(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}


	for j := 1; j <=  5; j ++ {
		scheduleTime := time.Now().Unix()
		for i := 1; i <= 10; i ++ {
			if i%2 == 0 {
				if err := SyncTackScheduleResult(1, 100, scheduleTime, true); err != nil {
					t.Fatal(err)
				}
			} else {
				if err := SyncTackScheduleResult(1, 100, scheduleTime, false); err != nil {
					t.Fatal(err)
				}
			}
		}
		time.Sleep(time.Second)
	}
}