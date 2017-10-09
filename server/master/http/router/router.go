package router

import (
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/http/handler"
)

func InitWorkerRouter(r *mux.Router) {
	sub := r.PathPrefix("/api/worker").Subrouter()
	sub.HandleFunc("/{wid}/ping", handler.PingHandler).Methods("GET")
	sub.HandleFunc("/{wid}", handler.RegisterWorkerHandler).Methods("POST")
}

func InitTaskRouter(r *mux.Router) {
	sub := r.PathPrefix("/api/task").Subrouter()

	//strategy and label are child resource of task
	sub.HandleFunc("", handler.CreateTaskHandler).Methods("POST")
	sub.HandleFunc("/{tid}", handler.DeleteTaskHandler).Methods("DELETE")
	sub.HandleFunc("/{tid}", handler.UpdateTaskHandler).Methods("PUT")
	sub.HandleFunc("/{tid}", handler.GetTaskHandler).Methods("GET")

	sub.HandleFunc("/{tid}/strategy", handler.AuthTaskMid(handler.CreateStrategyHandler)).Methods("POST")
	sub.HandleFunc("/{tid}/strategy", handler.AuthTaskMid(handler.DeleteStrategyHandler)).Methods("DELETE")
	sub.HandleFunc("/{tid}/strategy", handler.AuthTaskMid(handler.UpdateStrategyHandler)).Methods("PUT")
	sub.HandleFunc("/{tid}/strategy", handler.AuthTaskMid(handler.GetStrategyHandler)).Methods("GET")

	sub.HandleFunc("/{tid}/label", handler.EditLabelHandler).Methods("PUT")
}
