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
)

func EditLabelHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	tid := mux.Vars(r)["tid"]

	data, err := ioutil.ReadAll(io.LimitReader(r.Body, 200))
	if err != nil {
		log.Printf("update task(id=%s) label data(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	var label types.Label
	if err := json.Unmarshal(data, &label); err != nil {
		log.Printf("parse task(id=%s) label data(%s) err %v\n", tid, string(data), err)
		message.Error(w, err)
		return
	}

	if err := model.EditLabel(tid, &label); err != nil {
		log.Printf("update task(id=%s) label(%v) err %v\n", tid, label, err)
		message.Error(w, err)
	} else {
		message.Success(w)
	}
	return
}
