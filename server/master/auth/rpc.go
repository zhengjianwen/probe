package auth

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

var uic *Clients
var cloud *Clients

func initRpc(uicAddrs, cloudAddrs []string) {
	uic = NewClients(uicAddrs)
	cloud = NewClients(cloudAddrs)
}

func CallUIC(method string, args, reply interface{}) error {
	return uic.Call(method, args, reply)
}

func CallCloud(method string, args, reply interface{}) error {
	return cloud.Call(method, args, reply)
}

func NewClient(network, address string, timeout time.Duration) (*rpc.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
