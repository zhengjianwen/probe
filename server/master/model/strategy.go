package model

import (
	"github.com/rongyungo/probe/server/master/types"
	"gopkg.in/mgo.v2/bson"
)

func CreateStrategy(s *types.Strategy) error {
	c, close := getStrategyC()
	defer close()

	return c.Insert(s)
}

func DeleteStrategy(tid string) error {
	c, close := getStrategyC()
	defer close()

	return c.Remove(bson.M{"taskId": tid})
}

func UpdateStrategy(s *types.Strategy) error {
	c, close := getStrategyC()
	defer close()

	return c.Update(bson.M{"taskId": s.TaskId}, s)
}

func GetStrategy(tid string) (*types.Strategy, error) {
	c, close := getStrategyC()
	defer close()

	var s types.Strategy
	if err := c.Find(bson.M{"taskId": tid}).One(&s); err != nil {
		return nil, err
	}

	return &s, nil
}

func GetStrategyByTids(tids []string) ([]*types.Strategy, error) {
	c, close := getStrategyC()
	defer close()

	var l []*types.Strategy
	if err := c.Find(bson.M{"taskid": bson.M{"$in": tids}}).All(&l); err != nil {
		return nil, err
	}

	return l, nil
}
