package grpc

import (
	"github.com/rongyungo/probe/server/apm"
	pb "github.com/rongyungo/probe/server/proto"
	"log"
	"github.com/rongyungo/probe/server/master/model"
)

func handleResult(msg *pb.Topic) {
	if err := model.HandleTaskResult(msg.WorkerId, msg.Result); err != nil {
		log.Printf("server storage hand task result err %v\n", err)
	}

	if err := apm.Push(msg.Result); err != nil {
		log.Printf("server push task result to apm err %v\n", err)
	}
}
