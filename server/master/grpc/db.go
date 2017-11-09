package grpc

import (
	"github.com/rongyungo/probe/server/apm"
	"github.com/rongyungo/probe/server/master/model"
	pb "github.com/rongyungo/probe/server/proto"
	"log"
)

func handleResult(msg *pb.Topic) {
	model.CoverSnapShotM(msg.GetResult().GetType().String(), msg.GetResult().GetTaskId(), msg.WorkerId, msg.GetResult().GetDelayMs())

	if err := model.HandleTaskResult(msg.Result); err != nil {
		log.Printf("server storage hand task result err %v\n", err)
	}

	if err := apm.PushWorker(msg.WorkerId, msg.Result); err != nil {
		log.Printf("server push task result to apm err %v\n", err)
	}
}
