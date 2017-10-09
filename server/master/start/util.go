package start

import (
	"time"
)

func waitFn(sec time.Duration, fn func() error) error {
	c := make(chan error, 1)

	go func() {
		c <- fn()
	}()

	for {
		select {
		case <-time.Tick(sec):
			return nil
		case err := <-c:
			return err
		}
	}
}
