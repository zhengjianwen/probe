package types

import sqlutil "github.com/rongyungo/probe/util/sql"

type StartMasterConfig struct {
	HttpListeningAddress string
	GRpcListeningAddress string
	DbCfg                *sqlutil.DatabaseConfig
}
