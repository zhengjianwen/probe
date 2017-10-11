package handler

import (
	"github.com/1851616111/util/message"
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/model"
	"log"
	"net/http"
	"strconv"
)

func AuthTaskMid(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tid, ttp := mux.Vars(r)["tid"], mux.Vars(r)["ttp"]
		id, err := strconv.ParseInt(tid, 10, 64)
		if err != nil {
			log.Printf("auth middlerwre: get task id=%s invalide %v\n", tid, err)
			message.Error(w, err)
		}
		if _, err := model.GetTask(ttp, id); err != nil {
			log.Printf("auth middlerwre: get task id=%s err %v\n", tid, err)
			message.Error(w, err)
			return
		}

		fn(w, r)
	}
}
