package handler

import (
	"encoding/json"
	"github.com/1851616111/util/message"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"strconv"
)

func ReporterHandler(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)

	//uuid := vars["wid"]
	defer r.Body.Close()

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("accept worker' reporter data(%s) err %v\n", string(data), err)
		return
	}

	var reporter struct {
		Name string
	}

	if err := json.Unmarshal(data, &reporter); err != nil {
		log.Printf("parse worker's reporter data(%s) err %v\n", string(data), err)
		message.Error(w, err)
		return
	}

	message.Success(w)
}

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
