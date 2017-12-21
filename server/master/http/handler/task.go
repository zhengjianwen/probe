package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/1851616111/util/message"
	"github.com/rongyungo/probe/server/apm"
	cap "github.com/rongyungo/probe/server/img-cap"
	"github.com/rongyungo/probe/server/master/auth"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	errutil "github.com/rongyungo/probe/util/errors"
)

func GetTaskWorkerSnapShotHandler(w http.ResponseWriter, r *http.Request) {
	tType := mux.Vars(r)["ttp"]
	tidStr := mux.Vars(r)["tid"]
	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		message.Error(w, errutil.ErrTaskIdInvalid)
		return
	}

	if tType != "http" {
		message.NotFoundError(w)
		return
	}

	taskSS, ok := model.HttpSnapShotMapping[tid]
	if !ok {
		message.NotFoundError(w)
		return
	}

	message.SuccessI(w, taskSS)
	return
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ttp := mux.Vars(r)["ttp"]
	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)
	userID := r.Context().Value(auth.CONTEXT_KEY_USER).(int64)

	form, err := ParseForm(r.Body, ttp)
	if err != nil {
		log.Printf("creating task, parse form err %v\n", err)
		message.Error(w, err)
		return
	}
	if err := form.Validate(); err != nil {
		log.Printf("validate create task request err %v\n", err)
		message.Error(w, err)
		return
	}

	form.TaskObj.SetOrgId(orgId)

	if nodeId := form.TaskObj.GetNodeId(); nodeId > 0 {
		if ok, err := auth.CanWriteNode(userID, nodeId); err != nil {
			message.Error(w, err)
			return
		} else if !ok {
			message.Error(w, errors.New("no privileged"))
			return
		}
	}

	if _, err := model.CreateTask(form.TaskObj); err != nil {
		message.Error(w, err)
		return
	}

	imageName := fmt.Sprintf("task_url_%d.png", form.TaskObj.GetId())
	reqImageName := cap.GetReqImgName(imageName)
	form.TaskObj.SetWebImage(reqImageName)

	go cap.Cap(form.TaskObj.GetUrl(), cap.GetLocalImgName(imageName))

	ruleIds, err := apm.CreateRule(form)
	if err != nil {
		log.Printf("create task, regist apm rules error %v\n", err)
		message.Error(w, err)
		return
	}

	for _, ruleId := range ruleIds {
		form.TaskObj.AddRuleId(ruleId)
	}

	if err := model.UpdateTask(orgId, form.TaskObj.GetId(), form.TaskObj); err != nil {
		message.Error(w, err)
		return
	}
	message.Success(w)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	tidStr, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]

	userID := r.Context().Value(auth.CONTEXT_KEY_USER).(int64)
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

	ti, ok := task.(interface {
		GetNodeId() int64
	})
	if !ok {
		message.Error(w, errors.New("Server Inter Error"))
		return
	}

	if nodeId := ti.GetNodeId(); nodeId > 0 {
		if ok, err := auth.CanWriteNode(userID, nodeId); err != nil {
			message.Error(w, err)
			return
		} else if !ok {
			message.Error(w, errors.New("no privileged"))
			return
		}
	}

	if err := model.UpdateTask(orgId, tid, task); err != nil {
		message.Error(w, err)
	} else {
		//sc.UpdateTask(&task)
		message.Success(w)
	}
}

func OptTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	tidStr, _, statusStr := vars["tid"], vars["ttp"], vars["status"]

	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	var status bool
	if statusStr == "start" {
		status = true
	}
	var tp string = "http"
	if err := model.UpdateTaskStatus(orgId, tid, tp, status); err != nil {
		message.Error(w, err)
		return
	}

	message.Success(w)
}

func ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var nid int
	ttp, nidStr := mux.Vars(r)["ttp"], r.FormValue("nid")
	if len(nidStr) > 0 {
		var err error
		nid, err = strconv.Atoi(nidStr)
		if err != nil {
			message.Error(w, err)
			return
		}
	}

	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	if tasks, err := model.GetOrgTask(orgId, int64(nid), ttp); err != nil {
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

	taskI, err := model.GetTask(orgId, tid, ttp)
	if err != nil {
		message.Error(w, err)
		return
	}

	task, ok := taskI.(interface {
		GetRuleIds() []int64
	})
	if !ok {
		message.Error(w, err)
		return
	}

	if len(task.GetRuleIds()) > 0 {
		if err := apm.DeleteRules(orgId, task.GetRuleIds()...); err != nil {
			message.Error(w, err)
			return
		}
	}

	if err := model.DeleteTask(orgId, tid, ttp); err != nil {
		log.Printf("delete task(%s) err %v\n", tidStr, err)
		message.Error(w, err)
		return
	}

	//sc.DelTask(tidStr)
	message.Success(w)
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

func TaskOptRuleHandler(w http.ResponseWriter, r *http.Request) {
	ttp, tidStr, ruleIdStr, opt := mux.Vars(r)["ttp"], mux.Vars(r)["tid"], mux.Vars(r)["rid"], mux.Vars(r)["opt"]
	if opt != "bind" && opt != "unbind" {
		message.Error(w, errutil.ParamNotFoundErr("operate"))
		return
	}
	if len(tidStr) == 0 {
		message.Error(w, errutil.ParamNotFoundErr("task id"))
		return
	}
	if len(ttp) == 0 {
		message.Error(w, errutil.ParamNotFoundErr("task type"))
		return
	}
	if len(ruleIdStr) == 0 {
		message.Error(w, errutil.ParamNotFoundErr("rule id"))
		return
	}

	tid, err := strconv.ParseInt(tidStr, 10, 64)
	if err != nil {
		log.Printf("auth middlerwre: get task tid=%s invalide %v\n", tidStr, err)
		message.Error(w, err)
		return
	}

	ruleId, err := strconv.ParseInt(ruleIdStr, 10, 64)
	if err != nil {
		log.Printf("auth middlerwre: get rule id = %s invalide %v\n", tidStr, err)
		message.Error(w, err)
		return
	}

	orgId := r.Context().Value(auth.CONTEXT_KEY_ORG_ID).(int64)

	tk, err := model.GetTask(orgId, tid, ttp)
	if err != nil {
		message.Error(w, err)
		return
	}

	switch opt {
	case "bind":
		task, ok := tk.(interface {
			AddRuleId(int64)
		})
		if !ok {
			message.Error(w, errors.New("task type not satisfied"))
			return
		}
		task.AddRuleId(ruleId)
		err = model.UpdateTaskRuleId1(orgId, tid, task)

	case "unbind":
		task, ok := tk.(interface {
			RemoveRuleId(int64)
			GetRuleIds() []int64
		})
		if !ok {
			message.Error(w, errors.New("task type not satisfied"))
			return
		}
		task.RemoveRuleId(ruleId)
		err = model.UpdateTaskRuleId2(orgId, tid, ttp, task)
	}

	if err != nil {
		log.Printf("update task column rule_ids err %v\n", err)
		message.Error(w, err)
	} else {
		message.Success(w)
	}
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

func ParseForm(req io.Reader, tp string) (*types.CreateTaskForm, error) {
	data, err := ioutil.ReadAll(req)
	if err != nil {
		return nil, fmt.Errorf("read create task's request data err %v", err)
	}

	//创建task内容容器
	obj, ok := model.NewTaskPtr(tp).(types.CreateTaskI)
	if !ok {
		return nil, errors.New("unknown param task type")
	}

	//创建task业务容器
	form := types.CreateTaskForm{
		TaskObj: obj,
	}

	return &form, json.Unmarshal(data, &form)
}
