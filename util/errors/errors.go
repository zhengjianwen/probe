package util

import "errors"

var ErrTaskPeriodTooLess error = errors.New("task period too less")
var ErrUnSupportTaskType error = errors.New("task type un support")
var ErrTaskIdNotFound error = errors.New("task id not found")
var ErrTaskTypeMappingNotFound error = errors.New("task to type  mapping not found")
