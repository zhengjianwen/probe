package model

import (
	"github.com/rongyungo/probe/server/master/types"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func CreateTask(tk *types.Task) (string, error) {
	c, close := getTaskC()
	defer close()

	tk.CreateComplete()

	return tk.Id.Hex(), c.Insert(tk)
}

func DeleteTask(tid string) error {
	c, close := getTaskC()
	defer close()

	return c.RemoveId(tid)
}

func UpdateTask(tk *types.Task) error {
	c, close := getTaskC()
	defer close()
	tk.UpdateTime = time.Now().Unix()

	return c.UpdateId(tk.Id, tk)
}

func GetTask(tid string) (*types.Task, error) {
	c, close := getTaskC()
	defer close()

	var tk types.Task
	if err := c.FindId(bson.ObjectIdHex(tid)).One(&tk); err != nil {
		return nil, err
	}

	return &tk, nil
}

//@param inSec: the tasks should be executed in inSec second among the workers
//1. normal condition: time.Now() >= task.updateTime && time.Now() < (task.UpdateTime + task.PeriodSec)
//2. service down cause condition: time.Now() >= task.UpdateTime + task.PeriodSec
func GetAllTasks() ([]*types.Task, error) {
	c, close := getTaskC()
	defer close()

	var tasks []*types.Task
	if err := c.Find(nil).All(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetScheduleTasks() ([]*types.Task, error) {
	c, close := getTaskC()
	defer close()

	query := bson.M{
		"periodsec":    bson.M{"$gt": 0},
		"scheduletime": bson.M{"$lt": time.Now().Unix()},
	}

	var tasks []*types.Task
	if err := c.Find(query).All(&tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTaskTime(tk *types.Task) error {
	c, close := getTaskC()
	defer close()

	return c.UpdateId(tk.Id, bson.M{"$set": bson.M{
		"$executetime":  time.Now().Unix(),
		"$scheduleTime": tk.ScheduleTime + int64(tk.PeriodSec)},
	})
}

func CorrectScheduleTime() {
	c, close := getTaskC()
	defer close()

	now := time.Now().Unix()
	var tasks []*types.Task
	if err := c.Find(bson.M{"scheduleTime": bson.M{"$lt": now}}).All(&tasks); err != nil {
		log.Printf("scheduler: query correct schedule time err %v\n", err)
	}

	log.Printf("scheduler: query wrong scheduler time tasks: total: %d\n", len(tasks))

	var errN int
	for _, tk := range tasks {
		if err := c.UpdateId(tk.Id, bson.M{"$set": bson.M{"scheduleTime": now + int64(tk.PeriodSec)}}); err != nil {
			log.Printf("scheduler: correct schedule time err %v\n", err)
			errN++
		}
	}
	log.Printf("scheduler: query wrong scheduler time tasks: total: %d, error: %d\n", len(tasks), errN)
}
