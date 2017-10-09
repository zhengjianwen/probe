package scheduler

import (
	"github.com/ten-cloud/prober/server/master/model"
	"testing"
)

func TestInitTaskManager(t *testing.T) {
	if err := model.InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	err := InitTaskManager()
	if err != nil {
		t.Fatal(err)
	}

	StatTasks()
}
