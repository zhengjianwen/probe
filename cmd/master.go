package cmd

import (
	"fmt"
	"github.com/rongyungo/www/log"
	"github.com/spf13/cobra"
	"github.com/rongyungo/probe/server/master/start"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/rongyungo/probe/server/scheduler"
	"os"
)

var startMasterOptions startMasterOption

func init() {
	masterStartCmd.PersistentFlags().StringVarP(&startMasterOptions.gRpcListeningAddress, "grpc_listening", "", "127.0.0.1:9000", "master service grpc listening address")
	masterStartCmd.PersistentFlags().StringVarP(&startMasterOptions.httpListeningAddress, "http_listening", "", "127.0.0.1:9100", "master service http listening address")
	masterStartCmd.PersistentFlags().StringVarP(&startMasterOptions.databaseAddress, "database", "", "127.0.0.1:27017", "master service database address")
}

var masterCmd = &cobra.Command{
	Use:   "master",
	Short: "master operations command",
	Long:  `Distributed probe framework 's master command, this is prober's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

var masterStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start a master server command",
	Long:  `Distributed probe framework 's start master command, this is prober's`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := startMasterOptions.validate(); err != nil {
			fmt.Printf("master config validate err %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mCfg := types.StartMasterConfig{
			GRpcListeningAddress: startMasterOptions.gRpcListeningAddress,
			HttpListeningAddress: startMasterOptions.httpListeningAddress,
			DataBaseAddress:      startMasterOptions.databaseAddress,
		}
		rCfg := scheduler.RunConfig{
			DataBaseAddr: startMasterOptions.databaseAddress,
		}
		log.Printf("start all with config %#v\n", mCfg)
		if err := start.StartAll(&mCfg, &rCfg); err != nil {
			log.Errorf("start all fail %v\n", err)
			os.Exit(1)
		}
	},
}
