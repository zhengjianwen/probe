package model

import (
	"github.com/rongyungo/probe/server/master/types"
	"log"
)

var s types.Strategy

func CreateStrategy(s *types.Strategy) error {
	_, err := Orm.Insert(s)
	return err
}

func DeleteStrategy(tid int64) error {
	_, err := Orm.Id(tid).Delete(s)
	return err
}

func UpdateStrategy(s *types.Strategy) error {
	n, err := Orm.Id(s.TaskId).Count(s)
	if err != nil {
		return err
	}
	if n == 0 || n > 1 {
		log.Printf("invalid stragety tasi Idk conflict")
	}

	_, err = Orm.Id(s.TaskId).Update(s)

	return err
}

func GetStrategy(tid int64) (*types.Strategy, error) {
	s := types.Strategy{}
	err := Orm.Id(tid).Find(&s)
	if err != nil {
		return nil, err
	}
	return &s, nil
}
