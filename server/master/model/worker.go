package model

import (
	"github.com/ten-cloud/prober/server/master/types"
	"gopkg.in/mgo.v2/bson"
)

func RegisterWorker(worker *types.Worker) error {
	s, closeFn := getWorkerC()
	defer closeFn()

	_, err := s.Upsert(bson.M{"id": worker.ID}, worker)
	return err
}
