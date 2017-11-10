package model

import (
	"encoding/json"
	"fmt"
	"github.com/rongyungo/probe/server/master/types"
	"testing"
)

func Test_AdminEditWorker(t *testing.T) {
	//sql.DefaultDatabaseCfg.Host = "192.168.99.100"
	//if err := InitMySQL(&sql.DefaultDatabaseCfg); err != nil {
	//	t.Fatal(err)
	//}

	wk := &types.Worker{
		Country: "中国", Province: "", City: "深圳", Operator: "联通",
		Label: types.Label{
			Other: map[string]interface{}{"Location": []float32{114.07, 22.62}},
		},
	}
	//if err := AdminEditWorker(1, wk); err != nil {
	//	t.Fatal(err)
	//}

	data, _ := json.Marshal(wk)
	fmt.Printf("%s", string(data))
}
