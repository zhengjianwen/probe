package handler

import (
	"encoding/json"
	"github.com/1851616111/util/message"
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CreateStrategyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	tid := mux.Vars(r)["tid"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("create task(Id=%s) strategy(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	var strategy types.Strategy
	if err := json.Unmarshal(data, &strategy); err != nil {
		log.Printf("parse task(Id=%s) strategy(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	if err := strategy.Validate(); err != nil {
		log.Printf("validate new task(Id=%s) strategy(%v) err %v\n", tid, strategy, err)
		message.Error(w, err)
		return
	}

	strategy.TaskId = id
	if err := model.CreateStrategy(&strategy); err != nil {
		log.Printf("create task(Id=%s) strategy(%v) err %v\n", tid, strategy, err)
		message.Error(w, err)
	} else {
		message.Success(w)
	}
	return
}

func DeleteStrategyHandler(w http.ResponseWriter, r *http.Request) {
	tid := mux.Vars(r)["tid"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	if err := model.DeleteStrategy(id); err != nil {
		log.Printf("delete task(id=%s) strategy err %v\n", tid, err)
		message.Error(w, err)
	} else {
		message.Success(w)
	}
}

func UpdateStrategyHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	tid := mux.Vars(r)["tid"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("update task(id=%s) strategy data(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	var strategy types.Strategy
	if err := json.Unmarshal(data, &strategy); err != nil {
		log.Printf("parse task(id=%s) strategy data(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	if err := strategy.Validate(); err != nil {
		log.Printf("validate task(id=%s) updated strategy(%v) err %v\n", tid, strategy, err)
		message.Error(w, err)
		return
	}
	strategy.TaskId = id

	if err := model.UpdateStrategy(&strategy); err != nil {
		log.Printf("update task(id=%s) strategy(%v) err %v\n", tid, strategy, err)
		message.Error(w, err)
	} else {
		message.Success(w)
	}
	return
}

func GetStrategyHandler(w http.ResponseWriter, r *http.Request) {
	tid := mux.Vars(r)["tid"]
	id, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	if stg, err := model.GetStrategy(id); err != nil {
		log.Printf("get task(id=%s) strategy err %v\n", tid, err)
		message.Error(w, err)
	} else {
		message.SuccessI(w, stg)
	}
}
