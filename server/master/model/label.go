package model

import (
	"github.com/ten-cloud/prober/server/master/types"
	"gopkg.in/mgo.v2/bson"
)

func EditLabel(tid string, l *types.Label) error {
	c, close := getLabelC()
	defer close()

	_, err := c.Upsert(bson.M{"taskId": tid}, l)
	return err
}
