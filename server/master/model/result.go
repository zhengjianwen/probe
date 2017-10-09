package model

import (
	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
	"gopkg.in/mgo.v2/bson"
)

func InsertTaskResult(r *pb.TaskResult) error {
	fmt.Printf("-------->%#v\n", *r)
	c, close := getResultC()
	defer close()

	n, err := c.Find(bson.M{"taskid": r.TaskId}).Count()
	if err != nil {
		return err
	}

	target := bson.M{"error": r.Error, "errorcode": r.ErrorCode,
		"delayms": r.DelayMs, "success": r.Success, "startms": r.StartMs}

	if r.Ping != nil {
		target["ping"] = r.Ping
	}
	if r.Traceroute != nil {
		target["traceroute"] = r.Traceroute
	}
	if n == 0 {
		return c.Insert(bson.M{"taskid": r.TaskId, "taskList": []bson.M{target}})
	} else {
		return c.Update(bson.M{"taskid": r.TaskId}, bson.M{"$push": bson.M{"taskList": target}})
	}
}
