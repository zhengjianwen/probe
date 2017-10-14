package types

var (
	Worker_Status_New = "new"
)

type Worker struct {
	Id             int64
	Status         string
	StartTimestamp int64
	Country        string //国家
	Province       string //省
	City           string //市
	Operator       string //运营商
	Label          Label  `xorm:"json"`
}
