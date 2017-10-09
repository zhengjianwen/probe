package types

var (
	Worker_Status_New = "new"
)

type Worker struct {
	ID             string `json:"id", bson:"_id"`
	Status         string `json:"status", bson:"status"`
	StartTimestamp int64  `json:"startTimestamp", bson:"st"`
	Label          Label  `json:"label", bson:"label"`
}
