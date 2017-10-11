package handler

import (
	"encoding/json"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"log"
)

func createHttpTask(data []byte) (int64, error) {
	var task types.Task_Http
	if err := json.Unmarshal(data, &task); err != nil {
		log.Printf("parse http task data(%s) err %v\n", string(data), err)
		return 0, err
	}
	if err := task.Validate(); err != nil {
		log.Printf("validate new http task(%v) err %v\n", task, err)
		return 0, err
	}

	id, err := model.CreateTask_Http(&task)
	if err != nil {
		return 0, err
	}

	return id, err

	//scheduler.AddTask(&task)
}
