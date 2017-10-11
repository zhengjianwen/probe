package scheduler

import (
	"github.com/rongyungo/probe/server/master/model"
	"testing"
)

func TestNewTaskManager(t *testing.T) {
	if err := model.InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	NewTaskManager()
}
