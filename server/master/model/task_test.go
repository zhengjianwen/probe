package model

import (
	"fmt"
	"github.com/rongyungo/probe/server/master/types"
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/util/sql"
	"testing"
)

func TestCreateTask(t *testing.T) {
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
	fmt.Println(id)
}

//func TestCreateTask_Http(t *testing.T) {
//	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
//	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task_Http{
//		HttpSpec: pb.HttpSpec{
//			Url:       "http://www.baidu.com",
//			Method:    pb.HttpSpec_GET,
//			Header:    map[string]string{"k1": "v1", "k2": "v2"},
//			BasicAuth: &pb.HttpSpecBasicAuth{"michael", "123456"},
//		},
//		BasicInfo: pb.BasicInfo{
//			PeriodSec: 60,
//		},
//	}
//
//	id, err := CreateTask_Http(&task)
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(id)
//}
//
//func TestGetTask_Http(t *testing.T) {
//	sql.DefaultDatabaseCfg.Host = "192.168.99.100"
//	if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
//		t.Fatal(err)
//	}
//	th, err := GetTask_Http(1)
//	if err != nil && err != errutil.ErrTaskIdNotFound {
//		t.Fatal(err)
//	}
//
//	fmt.Printf("%#v\n", *th)
//}

//
//func Test_CreateTask_Http(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		HttpSpec: &pb.Task_Http{
//			Url:    "http://www.baidu.com",
//			Method: pb.Task_Http_GET,
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//

//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_CreateTask_Ping(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		PingSpec: &pb.Task_Ping{
//			Destination: "www.google.com",
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_CreateTask_TraceRoute(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		TraceRouteSpec: &pb.Task_Traceroute{
//			Destination: "www.google.com",
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_CreateTask_Tcp(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		TcpSpec: &pb.Task_Tcp{
//			Host: "127.0.0.1",
//			Port: 80,
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_CreateTask_Udp(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		UdpSpec: &pb.Task_Udp{
//			Host: "127.0.0.1",
//			Port: 80,
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_CreateTask_Ftp(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	task := types.Task{
//		FtpSpec: &pb.Task_Ftp{
//			Host:   "127.0.0.1",
//			Port:   80,
//			IfAuth: true,
//			Auth: &pb.Task_FtpAuth{
//				User:   "a",
//				Passwd: "b",
//			},
//		},
//		PeriodSec: 30,
//	}
//
//	if id, err := CreateTask(&task); err != nil {
//		t.Fatal(err)
//	} else {
//		fmt.Println(id)
//	}
//}
//
//func Test_GetAllTask(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	tList, err := GetAllTasks()
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	fmt.Println(len(tList))
//}
//
//func TestGetTask(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	_, err := GetTask("59ccfbcc4fa73e1344223c94")
//	if err != nil {
//		t.Fatal(err)
//	}
//}
