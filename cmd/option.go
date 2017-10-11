package cmd

import (
	"errors"
	"fmt"
	"github.com/1851616111/util/http"
)

type startMasterOption struct {
	gRpcListeningAddress string
	httpListeningAddress string
	databaseAddress      string
}

func (o startMasterOption) validate() error {
	if len(o.gRpcListeningAddress) == 0 {
		return errors.New("param gRpcListeningAddress not found")
	}

	if len(o.httpListeningAddress) == 0 {
		return errors.New("param httpListeningAddress not found")
	}

	return nil
}

type startWorkerOption struct {
	workerId int64
	pullSec  uint16

	masterGRpcAddresses []string
	masterHttpAddresses []string
}

func (o startWorkerOption) validate() error {

	if o.pullSec < 10 {
		return fmt.Errorf("param pullSec(%d) less than 10", o.pullSec)
	}

	for _, addr := range o.masterHttpAddresses {
		if err := validateMaster(o.workerId, addr); err != nil {
			return err
		}
	}
	return nil
}

func validateMaster(workerId int64, address string) error {
	spec := &http.HttpSpec{
		URL:         fmt.Sprintf("http://%s/api/worker/%d/ping", address, workerId),
		Method:      "GET",
		ContentType: http.ContentType_FORM,
	}
	fmt.Println(spec.URL)
	_, err := http.Send(spec)
	return err
}
