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
	errutil "github.com/rongyungo/probe/util/errors"
	"strings"
)

func GetTaskWorkerSnapShotHandler(w http.ResponseWriter, r *http.Request) {
	tType := strings.ToUpper(mux.Vars(r)["ttp"])
	tidStr := mux.Vars(r)["tid"]
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		message.Error(w, errutil.ErrTaskIdInvalid)
		return
	}

	taskSS, ok := model.TaskSnapShotMapping[tType][tid]
	if !ok {
		message.NotFoundError(w)
	}

	message.SuccessI(w, taskSS)
	return
}

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
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Printf("read task body data err %v\n", err)
		return nil, err
	}
	log.Printf("read task body data %s\n", string(data))

	target := model.GetTypeStructPtr(tp)
	if err := json.Unmarshal(data, target); err != nil {
		log.Printf("parse task body data err %v\n", err)
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
