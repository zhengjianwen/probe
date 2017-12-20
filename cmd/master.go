package cmd

import (
	"fmt"
	"github.com/rongyungo/probe/server/apm"
	"github.com/rongyungo/probe/server/master/auth"
	"github.com/rongyungo/probe/server/master/start"
	"github.com/rongyungo/probe/server/master/types"
	sqlutil "github.com/rongyungo/probe/util/sql"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var (
	startMasterOptions    startMasterOption
	DbCfg                 sqlutil.DatabaseConfig
	authCfg               auth.AuthConfig
	taskImageLocalHostDir string
	taskImageRequestPath  string
)

func init() {
	masterStartCmd.PersistentFlags().StringVarP(&startMasterOptions.gRpcListeningAddress, "grpc_listening", "", "0.0.0.0:9000", "master service grpc listening address")
	masterStartCmd.PersistentFlags().StringVarP(&startMasterOptions.httpListeningAddress, "http_listening", "", "0.0.0.0:9100", "master service http listening address")
	masterStartCmd.PersistentFlags().StringVarP(&DbCfg.Host, "host", "", "127.0.0.1", "master service database host")
	masterStartCmd.PersistentFlags().IntVarP(&DbCfg.Port, "port", "", 3306, "master service database port")
	masterStartCmd.PersistentFlags().StringVarP(&DbCfg.User, "user", "", "root", "master service database user name")
	masterStartCmd.PersistentFlags().StringVarP(&DbCfg.Password, "password", "", "123456", "master service database password")
	masterStartCmd.PersistentFlags().StringVarP(&DbCfg.DB, "instance", "", "probe", "master service database instance")

	masterStartCmd.PersistentFlags().IntVarP(&DbCfg.ConnMax, "max", "", 200, "master service database conn config")
	masterStartCmd.PersistentFlags().IntVarP(&DbCfg.ConnIdle, "idle", "", 100, "master service database conn config")

	masterStartCmd.PersistentFlags().StringVarP(&apm.Conf.Url, "apm-url", "", "http://www.opdeck.com", "master apm service url")
	masterStartCmd.PersistentFlags().StringVarP(&apm.Conf.Token, "apm-token", "", "ui49hfowlx0wkxoe,cjeaiqoei93ms8mx821kx", "master apm service token")

	masterStartCmd.PersistentFlags().StringVarP(&authCfg.CookieSecret, "cookie-secret", "c", "36005025221c6e1ff4bf9b255e49d356", "web cookie decode secret")
	masterStartCmd.PersistentFlags().StringVarP(&authCfg.CloudAddr, "svc-cloud", "", "127.0.0.1:3410", "rpc request address of cloud service")
	masterStartCmd.PersistentFlags().StringVarP(&authCfg.UicAddr, "scv-uic", "", "127.0.0.1:7772", "rpc request address of uic service")

	masterStartCmd.PersistentFlags().StringVarP(&taskImageLocalHostDir, "local-dir", "", "task_capture", "task capture store dir in host")
	masterStartCmd.PersistentFlags().StringVarP(&taskImageRequestPath, "req-path", "", "task_capture", "task capture store dir in host")

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

		if err := authCfg.Validate(); err != nil {
			fmt.Printf("auth config err %v\n", err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		mCfg := types.StartMasterConfig{
			GRpcListeningAddress: startMasterOptions.gRpcListeningAddress,
			HttpListeningAddress: startMasterOptions.httpListeningAddress,
			DbCfg:                &DbCfg,
			TaskImageLocalStoreDir: taskImageLocalHostDir,
		}

		log.Printf("start all with config %#v\n", mCfg)
		if err := start.RunAll(&mCfg, &DbCfg, &authCfg); err != nil {
			log.Printf("start all fail %v\n", err)
			os.Exit(1)
		}
	},
}
