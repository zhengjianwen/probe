package handler

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"
	"github.com/1851616111/util/message"
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	sc "github.com/rongyungo/probe/server/scheduler"
	"gopkg.in/mgo.v2/bson"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("create task data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	var task types.Task
	if err := json.Unmarshal(data, &task); err != nil {
		log.Printf("parse task data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	if err := task.Validate(); err != nil {
		log.Printf("validate new task(%v) err %v\n", task, err)
		message.Error(w, err)
		return
	}

	if id, err := model.CreateTask(&task); err != nil {
		log.Printf("create task(%v) err %v\n", task, err)
		message.Error(w, err)
	} else {
		sc.AddTask(&task)
		message.SuccessS(w, id)
	}
	return
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	tid := mux.Vars(r)["tid"]

	if err := model.DeleteTask(tid); err != nil {
		log.Printf("delete task(%s) err %v\n", tid, err)
		message.Error(w, err)
		return
	}

	sc.DelTask(tid)
	message.Success(w)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	tid := mux.Vars(r)["tid"]

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("update task(%s) data(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	var task types.Task
	if err := json.Unmarshal(data, &task); err != nil {
		log.Printf("parse task data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	if err := task.Validate(); err != nil {
		log.Printf("validate new task(%v) err %v\n", task, err)
		message.Error(w, err)
		return
	}
	task.Id = bson.ObjectId(tid)

	if err := model.UpdateTask(&task); err != nil {
		log.Printf("update task(%v) err %v\n", task, err)
		message.Error(w, err)
	} else {
		sc.UpdateTask(&task)
		message.Success(w)
	}
	return
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	tid := mux.Vars(r)["tid"]

	task, err := model.GetTask(tid)
	if err != nil {
		log.Printf("get task(%s) err %v\n", tid, err)
		message.Error(w, err)
	} else {
		message.SuccessI(w, task)
	}
	return
}
