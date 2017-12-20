package apm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/rongyungo/probe/server/master/types"
	"github.com/toolkits/net/httplib"
	"log"
)

type MapiUrlReply struct {
	Data struct {
		RuleId int64 `json:"ruleid"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func DeleteRules(orgId int64, ruleIds ...int64) error {
	for _, ruleId := range ruleIds {
		if err := DeleteRule(orgId, ruleId); err != nil {
			return err
		}
	}
	return nil
}
func DeleteRule(orgId, ruleId int64) error {
	reqAddr := fmt.Sprintf("%s/mapi/rule/%d/%d/delete?token=%s", Conf.Url, orgId, ruleId, Conf.Token)
	rsp, err := httplib.PostJSON(reqAddr, nil)
	if err != nil {
		log.Printf("delete apm(%s) rule(%d) err %s\n", reqAddr, ruleId, string(rsp))
	}
	return err
}

func CreateRule(form *types.CreateTaskForm) ([]int64, error) {
	//func CreateRule(orgId, taskId int64, url string, notify string, rules ...types.Rule) ([]int64, error) {
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
			"endpoints": fmt.Sprintf("url-%d", form.TaskObj.GetId()),
			//"metric":      "url.error",
			"metric":      newRule.Metric,
			"tags":        "",
			"func":        "all(#1)",
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
