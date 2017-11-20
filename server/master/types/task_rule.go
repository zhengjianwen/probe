package types

import (
	pb "github.com/rongyungo/probe/server/proto"
	"errors"
	"fmt"
)

type CreateTaskI interface {
	Validate() error
	SetId(int64)
	GetId()int64
	SetOrgId(int64)
	GetOrgId() int64
	AddRuleId(ruleId int64)
	GetNodeId() int64
	GetUrl() string
	GetType() pb.TaskType
	GetCreateTime() int64
	SetWebImage(string)
	Complete()
}

type CreateTaskForm struct {
	TaskObj CreateTaskI
	Rules   []Rule
	TeamIds []int
}

func (c CreateTaskForm) Validate() error {
	if len(c.TeamIds) == 0 {
		return errors.New("param TeamIds not found")
	}

	return c.TaskObj.Validate()
}

func (c CreateTaskForm) GetTeamIdsStr() string {
	var s string
	for idx, teamId := range c.TeamIds {
		s += fmt.Sprintf("%d", teamId)
		if idx < len(c.TeamIds) - 1 {
			s += ","
		}
	}
	return s
}

type Rule struct {
	MaxStep    int
	Metric     string
	RunBegin   string
	RunEnd     string
	Op         string
	RightValue float64
}
