package model

import (
	"gopkg.in/mgo.v2"
)

const (
	_Database      string = "prober"
	_Worker_C      string = "worker"
	_Task_C        string = "task"
	_TASK_RESULT_C string = "result"
	_Strategy_C    string = "strategy"
	_Label_C       string = "label"
)

var db *mgo.Session

func InitDb(addr string) error {
	var err error
	db, err = mgo.Dial(addr)

	return err
}

func NewDb(addr string) (*mgo.Session, error) {
	return mgo.Dial(addr)
}

func getWorkerC() (*mgo.Collection, func()) {
	s := db.Copy()
	return s.DB(_Database).C(_Worker_C), func() { s.Close() }
}

func getTaskC() (*mgo.Collection, func()) {
	s := db.Copy()
	return s.DB(_Database).C(_Task_C), func() { s.Close() }
}

func getResultC() (*mgo.Collection, func()) {
	s := db.Copy()
	return s.DB(_Database).C(_TASK_RESULT_C), func() { s.Close() }
}

func getStrategyC() (*mgo.Collection, func()) {
	s := db.Copy()
	return s.DB(_Database).C(_Strategy_C), func() { s.Close() }
}

func getLabelC() (*mgo.Collection, func()) {
	s := db.Copy()
	return s.DB(_Database).C(_Label_C), func() { s.Close() }
}
