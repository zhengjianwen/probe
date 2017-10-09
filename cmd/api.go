package cmd

import (
	"fmt"

	"encoding/json"
	"errors"
	"github.com/1851616111/util/http"
	"github.com/ten-cloud/prober/server/master/types"
	"io/ioutil"
	"time"
)

func registerWorker(opt *startWorkerOption) error {
	var wk types.Worker
	wk.ID = opt.workerId
	wk.Status = types.Worker_Status_New
	wk.StartTimestamp = time.Now().Unix()

	s := http.HttpSpec{
		URL:         fmt.Sprintf("http://%s/api/worker/%s", opt.masterHttpAddresses[0], wk.ID),
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		BodyObject:  wk,
	}

	rsp, err := http.Send(&s)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	var m struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	json.Unmarshal(data, &m)

	if m.Code != 1000 {
		return errors.New(m.Message)
	}

	return err
}
