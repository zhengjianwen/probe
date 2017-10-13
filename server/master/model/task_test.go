package model

import (
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/util/sql"
	"testing"
)

func TestCreateTask_Http(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Http{
		HttpSpec: pb.HttpSpec{
			Url:       "http://www.baidu.com",
			Method:    pb.HttpSpec_GET,
			Header:    map[string]string{"k1": "v1", "k2": "v2"},
			BasicAuth: &pb.HttpSpecBasicAuth{"michael", "123456"},
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_Ping(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Ping{
		PingSpec: pb.PingSpec{
			Destination: "www.google.com",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_TraceRoute(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_TraceRoute{
		TraceRouteSpec: pb.TraceRouteSpec{
			Destination: "www.google.com",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_Tcp(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Tcp{
		TcpSpec: pb.TcpSpec{
			Host: "www.google.com",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_Udp(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Udp{
		UdpSpec: pb.UdpSpec{
			Host: "www.google.com",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_Ftp(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Ftp{
		FtpSpec: pb.FtpSpec{
			Host: "www.google.com",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}

func Test_CreateTask_Dns(t *testing.T) {
	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
		t.Fatal(err)
	}

	task := types.Task_Dns{
		DnsSpec: pb.DnsSpec{
			Domain: "www.google.com",
			Type: pb.DnsSpec_A,
			ServerDesigned: true,
			DNSServer: "114.114.114.114",
		},
		BasicInfo: pb.BasicInfo{
			PeriodSec: 30,
		},
	}

	id, err := CreateTask(&task)
	if err != nil {
		t.Fatal(err)
	}
	if id == 0 {
		t.Fatal()
	}
}