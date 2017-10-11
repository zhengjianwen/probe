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
	sub.HandleFunc("/{ttp}", handler.CreateTaskHandler).Methods("POST")
	sub.HandleFunc("/{ttp}/{tid}", handler.DeleteTaskHandler).Methods("DELETE")
	sub.HandleFunc("/{ttp}/{tid}", handler.UpdateTaskHandler).Methods("PUT")
	sub.HandleFunc("/{ttp}/{tid}", handler.GetTaskHandler).Methods("GET")

	sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.CreateStrategyHandler)).Methods("POST")
	sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.DeleteStrategyHandler)).Methods("DELETE")
	sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.UpdateStrategyHandler)).Methods("PUT")
	sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.GetStrategyHandler)).Methods("GET")

	sub.HandleFunc("/{tid}/label", handler.EditLabelHandler).Methods("PUT")
}
