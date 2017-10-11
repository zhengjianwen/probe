package model

import (
	pb "github.com/rongyungo/probe/server/proto"
)

func InsertTaskResult(r *pb.TaskResult) error {
	_, err := Orm.Insert(r)
	return err
}
