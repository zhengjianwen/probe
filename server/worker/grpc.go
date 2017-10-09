package worker

import (
	pb "github.com/rongyungo/probe/server/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func NewClient(serverAddr string) (pb.MasterWorkerClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	return pb.NewMasterWorkerClient(conn), nil
}

func Subscribe(ctx context.Context, addr string, workerId string, in chan *pb.TaskResult, out chan *pb.TaskInfo, errCh chan error) {
	now := time.Now().Unix()
	topic := pb.Topic{
		Type:       pb.Topic_CONNECT,
		WorkerId:   workerId,
		WorkerTime: now,
	}

	latestStartTime = now
	cli, err := NewClient(addr)
	if err != nil {
		errCh <- err
		return
	}
	stream, err := cli.Subscribe(ctx)
	if err != nil {
		errCh <- err
		return
	}

	connectErrCh := make(chan error, 2)
	go func() {
		if err := stream.Send(&topic); err != nil {
			connectErrCh <- err
			return
		}

		for {
			select {
			case <-time.Tick(time.Minute / 3):
				topic.Type = pb.Topic_HEALTH_REPORT
				topic.WorkerTime = time.Now().Unix()
				if err := stream.Send(&topic); err != nil {
					connectErrCh <- err
					return
				}

			case result := <-in:
				if err := stream.Send(&pb.Topic{
					Type:   pb.Topic_RESULT,
					Result: result,
				}); err != nil {
					connectErrCh <- err
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			errCh <- ctx.Err()
			return
		case err := <-connectErrCh:
			log.Printf("client connect server err %v\n", err)
			return
		default:
			task, err := stream.Recv()
			if err != nil {
				errCh <- err
				return
			}
			log.Printf("recv master message %v\n", task.TaskId)
			out <- task
		}
	}
}
