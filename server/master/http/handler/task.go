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

	"github.com/rongyungo/probe/server/master/auth"
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
	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	task, err := readBodyToTask(r.Body, ttp, orgId)
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

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ttp := mux.Vars(r)["ttp"]
	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	if tasks, err := model.GetOrgTask(orgId, ttp); err != nil {
		message.Error(w, err)
	} else {
		message.SuccessI(w, tasks)
	}
	return
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	tidStr, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}
	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	if err := model.DeleteTask(orgId, tid, ttp); err != nil {
		log.Printf("delete task(%s) err %v\n", tidStr, err)
		message.Error(w, err)
		return
	}

	//sc.DelTask(tidStr)
	message.Success(w)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	tidStr, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]

	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)
	task, err := readBodyToTask(r.Body, ttp, orgId)
	if err != nil {
		message.Error(w, err)
		return
	}

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	if err := model.UpdateTask(orgId, tid, task); err != nil {
		message.Error(w, err)
	} else {
		//sc.UpdateTask(&task)
		message.Success(w)
	}
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	ttp, tidStr := mux.Vars(r)["ttp"], mux.Vars(r)["tid"]
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		log.Printf("auth middlerwre: get task tid=%s invalide %v\n", tidStr, err)
		message.Error(w, err)
		return
	}
	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)
	defer r.Body.Close()

	task, err := model.GetTask(orgId, tid, ttp)
	if err != nil {
		log.Printf("get task(%s) err %v\n", tidStr, err)
		message.Error(w, err)
	} else {
		message.SuccessI(w, task)
	}
	return
}


func readBodyToTask(rc io.Reader, tp string, orgId int64) (interface{}, error) {
	data, err := ioutil.ReadAll(rc)
	if err != nil {
		log.Printf("read task body data err %v\n", err)
		return nil, err
	}
	log.Printf("read task body data %s\n", string(data))

	target := model.NewTaskPtr(tp)
	if err := json.Unmarshal(data, target); err != nil {
		log.Printf("parse task body data err %v\n", err)
		return nil, err
	}

	v, _ := target.(interface {
		Validate() error
		SetOrgId(int64)
	})

	if err := v.Validate(); err != nil {
		return nil, err
	}

	v.SetOrgId(orgId)

	return target, nil
}
