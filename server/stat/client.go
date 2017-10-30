package stat

import (
	"net/http"
	"encoding/json"
	"github.com/rongyungo/probe/server/master/types"
	"time"
)

var MasterAddr = "http://127.0.0.1:9100"

func ListWorkingWorkers() ([]types.Worker, error) {
	rsp, err := http.Get(MasterAddr + "/probe/worker?source=memory")
	if err != nil {
		return nil, err
	}

	var wks []types.Worker
	if err := json.NewDecoder(rsp.Body).Decode(&wks); err != nil {
		return nil, err
	}

	return wks, nil
}

var WorkingWorks []types.Worker
var WorkingWorksTime int64 = time.Now().Unix()

func ListWorkerWithCached() ([]types.Worker, error) {
	// 缓存10分钟的working workers
	now := time.Now().Unix()
	if len(WorkingWorks) > 0 && now - 600 < WorkingWorksTime {
		return WorkingWorks, nil
	}

	var err error
	WorkingWorks, err = ListWorkingWorkers()
	WorkingWorksTime = now

	return WorkingWorks, err
}