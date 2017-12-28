package cmd

import (
	"fmt"

	"encoding/json"
	"errors"
	"github.com/1851616111/util/http"
	"github.com/rongyungo/probe/server/master/types"
	"io/ioutil"
	"log"
	"time"
)

func registerWorker(opt *startWorkerOption) error {
	now := time.Now().Unix()
	wk := types.Worker{
		Id:              opt.Id,
		Password:        opt.Password,
		Status:          types.Worker_Status_New,
		StartTimestamp:  now,
		UpdateTimestamp: now,
	}

	s := http.HttpSpec{
		URL:         fmt.Sprintf("http://%s/probe/worker/%d", opt.masterHttpAddresses[0], wk.Id),
		Method:      "POST",
		ContentType: http.ContentType_JSON,
		BodyObject:  wk,
	}

	log.Printf("------ %#v\n", s)

	rsp, err := http.Send(&s)
	if err != nil {
		log.Printf("----------- %v\n", err)
		return err
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Printf("----------- %v\n", err)
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

	return nil
}
