package apm

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/1851616111/util/http"
	"github.com/rongyungo/probe/server/master/types"
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
	req := http.HttpSpec{
		Method: "POST",
		URL:    reqAddr,
	}

	rsp, err := http.Send(&req)
	if err != nil {
		log.Printf("delete apm(%s) rule(%d) err %d\n", reqAddr, ruleId, string(rsp.Status))
	}
	return err
}

func CreateRule(form *types.CreateTaskForm) ([]int64, error) {
	url := fmt.Sprintf(
		"%s/mapi/rule/%d/create?team_ids=%s&token=%s",
		Conf.Url,
		form.TaskObj.GetOrgId(),
		form.GetTeamIdsStr(),
		Conf.Token,
	)

	var retRuleIds []int64
	for _, newRule := range form.Rules {
		data := map[string]interface{}{
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

		req := http.HttpSpec{
			Method:      "POST",
			URL:         url,
			BodyObject:  data,
			ContentType: http.ContentType_JSON,
		}
		resp, err := http.Send(&req)
		if err != nil {
			return nil, err
		}

		var ret MapiUrlReply
		err = json.NewDecoder(resp.Body).Decode(&ret)
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
