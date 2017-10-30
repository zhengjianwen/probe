package util

import "errors"

var ErrTaskPeriodTooLess error = errors.New("task period too less")
var ErrUnSupportTaskType error = errors.New("task type un support")
var ErrTaskNotFound error = errors.New("task not found")
var ErrTaskTypeMappingNotFound error = errors.New("task to type  mapping not found")
var ErrWorkerIdEmpty error = errors.New("worker id empty")
var ErrWorkerNotFound error = errors.New("worker not found")
var ErrWorkerUnAuthorized error = errors.New("worker unauthorized")
var ErrTaskIdInvalid error = errors.New("task id invalid")
var ErrWorkerIdInvalid error = errors.New("worker id invalid")
