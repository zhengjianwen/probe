package worker

import "log"

func Start(c *StartConfig) error {
	log.Printf("start work config %v\n", *c)
	m := NewTaskManager()
	StartController(c, m)

	return nil
}
