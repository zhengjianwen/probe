package handler

import (
	"encoding/json"
	"github.com/1851616111/util/message"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/server/master/grpc"
	"strconv"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write(([]byte)("pong"))
}

func RegisterWorkerHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	wid := mux.Vars(r)["wid"]
	id, err := strconv.ParseInt(wid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	var worker types.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		message.Error(w, err)
		return
	}

	worker.Id = id
	if err := worker.Validate(); err != nil {
		message.Error(w, err)
		return
	}

	log.Printf("new worker request %v\n", worker)

	if err := model.RegisterWorker(&worker); err != nil {
		message.Error(w, err)
		return
	}

	message.Success(w)
}

func AdminEditWorkerHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	wid := mux.Vars(r)["wid"]
	id, err := strconv.ParseInt(wid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	var worker types.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		message.Error(w, err)
		return
	}

	if err := worker.Validate(); err != nil {
		message.Error(w, err)
		return
	}

	log.Printf("edit worker request %v\n", worker)
	if err := model.AdminEditWorker(id, &worker); err != nil {
		message.Error(w, err)
		return
	}

	message.Success(w)
}

func AdminDelWorkerHandler(w http.ResponseWriter, r *http.Request) {
	wid := mux.Vars(r)["wid"]
	id, err := strconv.ParseInt(wid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}

	log.Printf("del worker request id=%d\n", id)
	if err := model.AdminDelWorker(id); err != nil {
		message.Error(w, err)
		return
	}
	message.Success(w)
}

func ListWorkersHandler(w http.ResponseWriter, r *http.Request) {
	grpc.Master.CleanWorkerConn()
	ids := grpc.Master.GetWorkerIds()

	if wks, err := model.ListWorkers(ids...); err != nil {
		message.Error(w, err)
	} else {
		message.SuccessI(w, wks)
	}
	return
}

func GetWorkerHandler(w http.ResponseWriter, r *http.Request) {
	wid := mux.Vars(r)["wid"]
	id, err := strconv.ParseInt(wid, 10, 64)
	if err != nil {
		message.Error(w, err)
		return
	}
	if wk, err := model.GetWorkerById(id); err != nil {
		message.Error(w, err)
	} else {
		message.SuccessI(w, wk)
	}
	return
}
