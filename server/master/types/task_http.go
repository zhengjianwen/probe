package types

import (
	"github.com/ten-cloud/prober/server/master/util"
	pb "github.com/ten-cloud/prober/server/proto"
)

const (
	minUrlLen = 4 //a.cn
	maxUrlLen = 1024
)

type Task_Http_Spec struct {
	Url           string `json:"url"`
	BodyMatchText string `json:"matchText"`
	Timeout       uint32 `json:"timeout"` //ms
	StatusCode    uint16 `json:"statusCode"`
	Retry         uint8  `json:"retry"`
	Creator       int64  `json:"creator"`
}

func Validate(t *pb.Task_Http) error {
	if len(t.Url) < minUrlLen || len(t.Url) > maxUrlLen {
		return util.ParamLengthInvalid("url", minUrlLen, maxUrlLen)
	}
	return nil
}
