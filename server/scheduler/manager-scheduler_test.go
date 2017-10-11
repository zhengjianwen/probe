package scheduler

import (
	sqlutil "github.com/rongyungo/probe/util/sql"
	"testing"
)

func TestNewSchedulerManager(t *testing.T) {
	sqlutil.DefaultDatabaseCfg.Host = "192.168.99.100"
	if _, err := NewSchedulerManager("http", 60, &sqlutil.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}
}

func TestScheduleManager_Start(t *testing.T) {
	sqlutil.DefaultDatabaseCfg.Host = "192.168.99.100"
	m, err := NewSchedulerManager("http", 60, &sqlutil.DefaultDatabaseCfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := m.Start(); err != nil {
		t.Fatal(err)
	}

	select {}
}
