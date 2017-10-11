package start

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rongyungo/probe/server/master/grpc"
	"github.com/rongyungo/probe/server/master/http/router"
	"github.com/rongyungo/probe/server/master/model"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/server/scheduler"
	"github.com/rongyungo/probe/util/sql"
)

func StartAll(mCfg *types.StartMasterConfig, dbc *sql.DatabaseConfig) error {
	if err := model.InitMySQL(dbc); err != nil {
		return err
	}

	if err := StartScheduler(dbc); err != nil {
		return err
	}

	if err := StartMaster(mCfg); err != nil {
		return err
	}

	return nil
}

func StartMaster(cfg *types.StartMasterConfig) error {

	rt := mux.NewRouter()
	router.InitWorkerRouter(rt)
	router.InitTaskRouter(rt)

	log.Printf("starting grpc service, listening on %s port\n", cfg.GRpcListeningAddress)
	waitFn(time.Second*3, func() error {
		return grpc.StartServer(&grpc.StartConfig{
			ListeningAddress: cfg.GRpcListeningAddress,
		})
	})
	log.Printf("grpc service start success.")

	log.Printf("starting http service, listening on %s port\n", cfg.HttpListeningAddress)
	waitFn(time.Second*3, func() error {
		return http.ListenAndServe(cfg.HttpListeningAddress, rt)
	})
	log.Printf("http service start success.")

	select {}
	return nil
}

func StartScheduler(dbc *sql.DatabaseConfig) error {
	var m1, m2, m3, m4, m5, m6, m7 *scheduler.ScheduleManager
	var err error
	if m1, err = scheduler.NewSchedulerManager("http", 60, dbc); err != nil {
		return err
	}
	if m2, err = scheduler.NewSchedulerManager("dns", 60, dbc); err != nil {
		return err
	}
	if m3, err = scheduler.NewSchedulerManager("ftp", 60, dbc); err != nil {
		return err
	}
	if m4, err = scheduler.NewSchedulerManager("ping", 60, dbc); err != nil {
		return err
	}
	if m5, err = scheduler.NewSchedulerManager("tcp", 60, dbc); err != nil {
		return err
	}
	if m6, err = scheduler.NewSchedulerManager("udp", 60, dbc); err != nil {
		return err
	}
	if m7, err = scheduler.NewSchedulerManager("trace_route", 60, dbc); err != nil {
		return err
	}

	if err = m1.Start(); err != nil {
		return err
	}
	if err = m2.Start(); err != nil {
		return err
	}
	if err = m3.Start(); err != nil {
		return err
	}
	if err = m4.Start(); err != nil {
		return err
	}
	if err = m5.Start(); err != nil {
		return err
	}
	if err = m6.Start(); err != nil {
		return err
	}
	if err = m7.Start(); err != nil {
		return err
	}

	return nil
}
