package worker

import (
	"context"
	pb "github.com/rongyungo/probe/server/proto"
	"github.com/rongyungo/probe/server/worker/exec"
	"log"
	"time"
)

var (
	conErrCh            chan error = make(chan error, 10)
	latestStartTime     int64
	allowStartPeriodSec int64 = 60
)

func StartController(c *StartConfig) error {
	var (
		ctx      context.Context
		cancel   context.CancelFunc
		taskCh   chan *pb.Task       = make(chan *pb.Task, 100)
		resultCh chan *pb.TaskResult = make(chan *pb.TaskResult, 100)
	)

	conErrCh <- nil

	for {
		select {
		case err := <-conErrCh:
			if err != nil {
				log.Printf("worker subscribe goroutine enter err %v, will restart\n", err)
				cancel()
			}

			wait()
			log.Printf("worker prepare start subscribe goroutine\n")
			ctx, cancel = context.WithCancel(context.Background())
			go Subscribe(ctx, c.MasterGRpcs[0], c.WorkerId, resultCh, taskCh, conErrCh)

		case t := <-taskCh:
			log.Printf("prepare exec task(%s) \n", t.String())
			//m.AddTask(t)

			go func(task *pb.Task) {
				if res := exec.Execute(task); res != nil {
					log.Printf("exec %s task(%s) result %s\n",
						t.GetBasicInfo().GetType().String(), t.GetBasicInfo().GetId(), res.String())
					resultCh <- res
				}
			}(t)

		}
	}
	return nil
}

func wait() {
	startIntervalSec := time.Now().Unix() - latestStartTime
	if startIntervalSec <= allowStartPeriodSec {
		time.Sleep(time.Second * time.Duration(allowStartPeriodSec-startIntervalSec))
	}

	return
}
