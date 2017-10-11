package model

//
//import (
//	"fmt"
//	"github.com/rongyungo/probe/server/master/types"
//	pb "github.com/rongyungo/probe/server/proto"
//	"testing"
//)
//
//func TestCreateStrategy(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	//task := types.Task{
//	//	HttpSpec: &pb.Task_Http{
//	//		Url:    "http://www.baidu.com",
//	//		Method: pb.Task_Http_GET,
//	//	},
//	//	PeriodSec: 30,
//	//}
//
//	id, err := CreateTask(&task)
//	if err != nil {
//		t.Fatal(err)
//	}
//
//	if len(id) == 0 {
//		t.Fatal("create task not return taskId")
//	}
//
//	strategy := types.Strategy{
//		TaskId: id,
//		All:    true,
//	}
//
//	if err := CreateStrategy(&strategy); err != nil {
//		t.Fatal(err)
//	}
//}
//
//func TestGetStrategyByTids(t *testing.T) {
//	if err := InitDb("127.0.0.1:27017"); err != nil {
//		t.Fatal(err)
//	}
//
//	s, err := GetStrategyByTids([]string{"59c387624fa73e04583d3b06"})
//	if err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println(s)
//}
