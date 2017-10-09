package types

type StartMasterConfig struct {
	HttpListeningAddress string
	GRpcListeningAddress string
	DataBaseAddress      string
}
