package router

import (
	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/auth"
	"github.com/rongyungo/probe/server/master/http/handler"
)

func InitWorkerRouter(r *mux.Router) {
	sub := r.PathPrefix("/probe/worker").Subrouter()
	sub.HandleFunc("", handler.ListWorkersHandler).Methods("GET")
	sub.HandleFunc("/{wid}", handler.GetWorkerHandler).Methods("GET")
	sub.HandleFunc("/{wid}", handler.RegisterWorkerHandler).Methods("POST")
	sub.HandleFunc("/{wid}", handler.AuthAdmin(handler.AdminEditWorkerHandler)).Methods("PUT")
	sub.HandleFunc("/{wid}", handler.AuthAdmin(handler.AdminDelWorkerHandler)).Methods("DELETE")
	sub.HandleFunc("/{wid}/ping", handler.PingHandler).Methods("GET")
}

func InitTaskRouter(r *mux.Router) {
	sub := r.PathPrefix("/probe/task").Subrouter()

	sub.HandleFunc("/{ttp}/{tid}/snapshot", handler.GetTaskWorkerSnapShotHandler).Methods("GET")

	//strategy and label are child resource of task
	sub.HandleFunc("/org/{oid}/{ttp}", auth.AuthUser(handler.CreateTaskHandler)).Methods("POST")
	sub.HandleFunc("/org/{oid}/{ttp}", auth.AuthUser(handler.ListTasksHandler)).Methods("GET")
	sub.HandleFunc("/org/{oid}/{ttp}/task/{tid}", auth.AuthUser(handler.GetTaskHandler)).Methods("GET")
	sub.HandleFunc("/org/{oid}/{ttp}/task/{tid}", auth.AuthUser(handler.DeleteTaskHandler)).Methods("DELETE")
	sub.HandleFunc("/org/{oid}/{ttp}/task/{tid}", auth.AuthUser(handler.UpdateTaskHandler)).Methods("PUT")

	sub.HandleFunc("/org/{oid}/{ttp}/task/{tid}/{opt}/rule/{rid}", auth.AuthUser(handler.TaskOptRuleHandler)).Methods("POST")
	//sub.HandleFunc("/{oid}/{ttp}/{tid}/unbind/{rid}", auth.AuthUser(handler.TaskOptRuleHandler)).Methods("POST")

	//sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.CreateStrategyHandler)).Methods("POST")
	//sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.DeleteStrategyHandler)).Methods("DELETE")
	//sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.UpdateStrategyHandler)).Methods("PUT")
	//sub.HandleFunc("/{ttp}/{tid}/strategy", handler.AuthTaskMid(handler.GetStrategyHandler)).Methods("GET")

	sub.HandleFunc("/{tid}/label", handler.EditLabelHandler).Methods("PUT")
}
