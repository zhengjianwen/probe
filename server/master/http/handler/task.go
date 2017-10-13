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
	"strconv"
)

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ttp := mux.Vars(r)["ttp"]

	task, err := readBodyToTask(r.Body, ttp)
	if err != nil {
		message.Error(w, err)
		return
	}

	if id, err := model.CreateTask(task); err != nil {
		message.Error(w, err)
	} else {
		//	sc.AddTask(&task)
		message.SuccessS(w, fmt.Sprintf("%d", id))
	}
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
	defer r.Body.Close()
	tid, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	task, err := readBodyToTask(r.Body, ttp)
	if err != nil {
		message.Error(w, err)
		return
	}

	if err := model.UpdateTask(id, task); err != nil {
		message.Error(w, err)
	} else {
		//sc.UpdateTask(&task)
		message.Success(w)
	}
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

func readBodyToTask(rc io.Reader, tp string) (interface{}, error) {
	data, err := ioutil.ReadAll(io.LimitReader(rc, 200))
	if err != nil {
		return nil, err
	}

	target := model.GetTypeStructPtr(tp)
	if err := json.Unmarshal(data, target); err != nil {
		return nil, err
	}

	v, _ := target.(interface {
		Validate() error
	})

	if err := v.Validate(); err != nil {
		return nil, err
	}

	return target, nil
}
