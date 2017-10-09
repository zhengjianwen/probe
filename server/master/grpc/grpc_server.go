package grpc

import (
	"net"

	"github.com/rongyungo/probe/server/master/model"
	pb "github.com/rongyungo/probe/server/proto"
	"google.golang.org/grpc"
	"log"
)

func StartServer(c *StartConfig) error {
	lis, err := net.Listen("tcp", c.ListeningAddress)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	gServer := grpc.NewServer(opts...)
	if Master == nil {
		Master = NewMaster()
	}

	pb.RegisterMasterWorkerServer(gServer, (pb.MasterWorkerServer)(Master))

	log.Fatal(gServer.Serve(lis))
	return nil
}

func (m *master) Subscribe(stream pb.MasterWorker_SubscribeServer) error {
	var closeServeCh chan<- struct{}
	for {
		msg, err := stream.Recv()
		if err != nil {
			if closeServeCh != nil {
				closeServeCh <- struct{}{}
			}
			log.Printf("<<master recv worker's subscribe err %v, close subscribe>>\n", err)
			return err
		}

		switch msg.Type {
		case pb.Topic_CONNECT:
			closeServeCh = m.serveWorker(msg.WorkerId, stream)
		case pb.Topic_HEALTH_REPORT:
			m.healthCheck(msg.WorkerId)
		case pb.Topic_RESULT:
			if err := model.InsertTaskResult(msg.Result); err != nil {
				log.Printf("server store task result err %v\n", err)
			}
		}
		log.Printf("server recv message %#v\n", msg.String())
	}
	return nil
}

type StartConfig struct {
	ListeningAddress string
	SubPeriodSec     uint16 //worker sub period second
}
