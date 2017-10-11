package cmd

import (
	"fmt"
	"github.com/1851616111/util/rand"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/rongyungo/probe/server/worker"
)

var startWorkerOptions startWorkerOption = startWorkerOption{
	masterGRpcAddresses: []string{},
	masterHttpAddresses: []string{},
}

func init() {
	workerStartCmd.PersistentFlags().Uint16VarP(&startWorkerOptions.pullSec, "period", "p", 60, "worker service report period second")
	workerStartCmd.PersistentFlags().StringSliceVarP(&startWorkerOptions.masterHttpAddresses, "master_http", "", []string{"127.0.0.1:9100"}, "worker's master http addresses")
	workerStartCmd.PersistentFlags().StringSliceVarP(&startWorkerOptions.masterGRpcAddresses, "master_grpc", "", []string{"127.0.0.1:9000"}, "worker's master grpc ddresses")
	workerStartCmd.PersistentFlags().Int64VarP(&startWorkerOptions.workerId, "id", "i", int64(rand.Intn(100000)), "worker name that user assigned")
}

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "worker operations command",
	Long: `Distributed framework's worker command, this is prober's.
			worker accept scheduler's tasks and perform it.
			worker also need timely report to master about the system and task information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var workerStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a worker service command",
	Long:  `Distributed probe framework 's start worker command, this is prober's`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := startWorkerOptions.validate(); err != nil {
			log.Printf("start worker failed, cause %v \n", err)
			os.Exit(1)
		}

		log.Printf("validate master address %s ok.\n", startWorkerOptions.masterHttpAddresses)
		log.Printf("worker id is %s.\n", startWorkerOptions.workerId)

	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := registerWorker(&startWorkerOptions); err != nil {
			log.Printf("register worker failed, cause %v \n", err)
			os.Exit(1)
		}

		log.Printf("register worker %s success\n", startWorkerOptions.workerId)

		worker.Start(&worker.StartConfig{
			WorkerId:       startWorkerOptions.workerId,
			HealthCheckSec: startWorkerOptions.pullSec,
			MasterHttps:    startWorkerOptions.masterHttpAddresses,
			MasterGRpcs:    startWorkerOptions.masterGRpcAddresses,
		})
	},
}
