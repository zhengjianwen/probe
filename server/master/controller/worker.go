package controller

//import (
//	"time"
//	"github.com/rongyungo/probe/server/master/model"
//	"github.com/rongyungo/probe/server/master/grpc"
//
//)
//
////向数据库同步worker信息用于stat统计
//func StartWorkerController() {
//	tk := time.NewTicker(time.Second * time.Duration(5))
//	for {
//		select {
//		case <-tk.C:
//			grpc.Master.CleanWorkerConn()
//
//		}
//	}
//}
