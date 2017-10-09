package handler

import (
	"github.com/1851616111/util/message"
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"log"
	"net/http"
)

func AuthTaskMid(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tid := mux.Vars(r)["tid"]
		if _, err := model.GetTask(tid); err != nil {
			log.Printf("auth middlerwre: get task id=%s err %v\n", tid, err)
			message.Error(w, err)
			return
		}

		fn(w, r)
	}
}
