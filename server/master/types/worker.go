package types

var (
	Worker_Status_New = "new"
)

type Worker struct {
	Id             int64
	Status         string
	StartTimestamp int64
	Label          Label `xorm:"json"`
}
