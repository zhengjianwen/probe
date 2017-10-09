package start

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ten-cloud/prober/server/master/grpc"
	"github.com/ten-cloud/prober/server/master/http/router"
	"github.com/ten-cloud/prober/server/master/model"
	"github.com/ten-cloud/prober/server/master/types"
	"github.com/ten-cloud/prober/server/scheduler"
)

func StartAll(mCfg *types.StartMasterConfig, sCfg *scheduler.RunConfig) error {
	if err := model.InitDb(mCfg.DataBaseAddress); err != nil {
		return err
	}

	if err := scheduler.StartScheduler(sCfg); err != nil {
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
