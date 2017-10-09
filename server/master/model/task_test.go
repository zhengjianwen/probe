package model

import (
	"fmt"
	"github.com/ten-cloud/prober/server/master/types"
	pb "github.com/ten-cloud/prober/server/proto"
	"testing"
)

func Test_CreateTask_Http(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	task := types.Task{
		HttpSpec: &pb.Task_Http{
			Url:    "http://www.baidu.com",
			Method: pb.Task_Http_GET,
		},
		PeriodSec: 30,
	}

	if id, err := CreateTask(&task); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(id)
	}
}

func Test_CreateTask_Dns(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	task := types.Task{
		DnsSpec: &pb.Task_Dns{
			Domain:         "www.baidu.com",
			Type:           pb.Task_Dns_A,
			ServerDesigned: true,
			DNSServer:      "114.114.114.114:53",
		},
		PeriodSec: 30,
	}

	if id, err := CreateTask(&task); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(id)
	}
}

func Test_CreateTask_Ping(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	task := types.Task{
		PingSpec: &pb.Task_Ping{
			Destination: "www.google.com",
		},
		PeriodSec: 30,
	}

	if id, err := CreateTask(&task); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(id)
	}
}

func Test_CreateTask_TraceRoute(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	task := types.Task{
		TraceRouteSpec: &pb.Task_Traceroute{
			Destination: "www.google.com",
		},
		PeriodSec: 30,
	}

	if id, err := CreateTask(&task); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(id)
	}
}

func Test_GetAllTask(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	tList, err := GetAllTasks()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(tList))
}

func TestGetTask(t *testing.T) {
	if err := InitDb("127.0.0.1:27017"); err != nil {
		t.Fatal(err)
	}

	_, err := GetTask("59ccfbcc4fa73e1344223c94")
	if err != nil {
		t.Fatal(err)
	}
}
