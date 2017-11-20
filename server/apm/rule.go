package apm

import (
	"fmt"
	"errors"
	"encoding/json"
	"github.com/toolkits/net/httplib"
	"github.com/rongyungo/probe/server/master/types"
)

type MapiUrlReply struct {
	Data struct {
		RuleId int64 `json:"ruleid"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func ApmRuleRegister(form *types.CreateTaskForm) ([]int64, error) {
//func ApmRuleRegister(orgId, taskId int64, url string, notify string, rules ...types.Rule) ([]int64, error) {
	apmSvcAddr := fmt.Sprintf(
		"%s/mapi/rule/%d/create?team_ids=%s&token=%s",
		Conf.Url,
		form.TaskObj.GetOrgId(),
		form.GetTeamIdsStr(),
		Conf.Token,
	)

	var retRuleIds []int64
	for _, newRule := range form.Rules {
		args := map[string]interface{}{
			"endpoints":   fmt.Sprintf("url-%d", form.TaskObj.GetId()),
			//"metric":      "url.error",
			"metric":      newRule.Metric,
			"tags":        "",
			"func":       "all(#1)",
			"op":          newRule.Op,
			"right_value": newRule.RightValue,
			"max_step":    newRule.MaxStep,
			"priority":    0,
			"note":        "[URL: " + form.TaskObj.GetUrl() + "]",
			"run_begin":   newRule.RunBegin,
			"run_end":     newRule.RunEnd,
		}

		resp, err := httplib.PostJSON(apmSvcAddr, args)
		if err != nil {
			return nil, err
		}

		var ret MapiUrlReply
		err = json.Unmarshal(resp, &ret)
		if err != nil {
			return nil, err
		}

		if ret.Msg != "" {
			return nil, errors.New(fmt.Sprintf("internal server error: mapi error: %s", ret.Msg))
		}

		retRuleIds = append(retRuleIds, ret.Data.RuleId)
	}

	return retRuleIds, nil
}