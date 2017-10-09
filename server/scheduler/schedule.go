package scheduler

import (
	"github.com/ten-cloud/prober/server/master/types"
	"log"
)

func Schedule(ws []string, s *types.Strategy, t *types.Task, sendFn func(wid string, tk *types.Task) error, onSuccessFn func(tk *types.Task) error) error {
	for _, wid := range ws {
		log.Printf("schedule task(%s) to worker(%s)\n", t.Id.Hex(), wid)
		err := sendFn(wid, t)
		if err != nil {
			log.Printf("schedule task(%s) to worker(%s) err %v\n", t.Id.Hex(), wid, err)
			//statis
			continue
		}
		if err := onSuccessFn(t); err != nil {
			//statis
		}

	}

	return nil
}
