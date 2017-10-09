package scheduler

import (
	"github.com/ten-cloud/prober/server/master/grpc"
	"github.com/ten-cloud/prober/server/master/model"
	"log"
	"time"
)

func StartScheduler(c *RunConfig) error {
	if err := InitTaskManager(); err != nil {
		return err
	}

	go Run(c)

	StatTasks()
	return nil
}

func Run(c *RunConfig) {
	ctk := time.NewTicker(time.Second * time.Duration(60)) //scheduler time correct ticker
	tk := time.NewTicker(time.Second * time.Duration(5))
	for {
		select {
		case <-tk.C:
			tasks, err := model.GetScheduleTasks()
			if err != nil {
				log.Printf("scheduler get to schedule tasks err %v\n", err)
				continue
			}

			tasks = ReduceReplicatedTask(tasks)
			if len(tasks) > 0 {
				log.Printf("scheduler: query prepare schedule tasks: %d\n", len(tasks))
			}

			workerIds := grpc.Master.GetWorkerIds()

			for _, tk := range tasks {
				Schedule(workerIds, nil, tk, grpc.Master.SendTask, model.UpdateTaskTime)
			}

		case <-ctk.C:
			go model.CorrectScheduleTime()
		}
	}
}
