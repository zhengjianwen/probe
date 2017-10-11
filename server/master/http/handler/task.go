package handler

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"
	"fmt"
	"github.com/1851616111/util/message"
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"strconv"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	ttp := mux.Vars(r)["ttp"]
	defer r.Body.Close()

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("create task data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	var target interface{}
	switch ttp {
	case "http":
		target = &types.Task_Http{}
	case "dns":
		target = &types.Task_Dns{}
	case "ping":
		target = &types.Task_Ping{}
	case "trace_route":
		target = &types.Task_TraceRoute{}
	case "tcp":
		target = &types.Task_Tcp{}
	case "udp":
		target = &types.Task_Udp{}
	case "ftp":
		target = &types.Task_Ftp{}
	}

	v, _ := target.(interface {
		Validate() error
	})
	if err := v.Validate(); err != nil {
		message.Error(w, err)
		return
	}

	if err := json.Unmarshal(data, target); err != nil {
		log.Printf("parse http task data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	if id, err := model.CreateTask(target); err != nil {
		message.Error(w, err)
	} else {
		//	sc.AddTask(&task)
		message.SuccessS(w, fmt.Sprintf("%d", id))
	}

	return
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	tid, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	if err := model.DeleteTask(ttp, id); err != nil {
		log.Printf("delete task(%s) err %v\n", tid, err)
		message.Error(w, err)
		return
	}

	//sc.DelTask(tid)
	message.Success(w)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	//defer r.Body.Close()
	//
	//tid := mux.Vars(r)["tid"]
	//
	//data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	//if err != nil {
	//	log.Printf("update task(%s) data(%s) err %v\n", tid, string(data), err)
	//	message.Error(w, err)
	//	return
	//}
	//
	//var task types.Task
	//if err := json.Unmarshal(data, &task); err != nil {
	//	log.Printf("parse task data(%s) err %v\n", string(data), err)
	//	message.Error(w, err)
	//	return
	//}
	//
	//if err := task.Validate(); err != nil {
	//	log.Printf("validate new task(%v) err %v\n", task, err)
	//	message.Error(w, err)
	//	return
	//}
	//task.Id = bson.ObjectId(tid)
	//
	//if err := model.UpdateTask(&task); err != nil {
	//	log.Printf("update task(%v) err %v\n", task, err)
	//	message.Error(w, err)
	//} else {
	//	sc.UpdateTask(&task)
	//	message.Success(w)
	//}
	//return
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	ttp, tid := mux.Vars(r)["ttp"], mux.Vars(r)["tid"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		log.Printf("auth middlerwre: get task id=%s invalide %v\n", tid, err)
		message.Error(w, err)
	}
	defer r.Body.Close()

	task, err := model.GetTask(ttp, id)
	if err != nil {
		log.Printf("get task(%s) err %v\n", tid, err)
		message.Error(w, err)
	} else {
		message.SuccessI(w, task)
	}
	return
}
