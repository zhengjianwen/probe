package exec

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	pb "github.com/rongyungo/probe/server/proto"
	"time"
)

func ProbeFtp(t *pb.TaskInfo) *pb.TaskResult {
	var start int64 = time.Now().UnixNano()
	var res pb.TaskResult = pb.TaskResult{TaskId: t.TaskId, StartMs: start / 1e6}

	con, err := ftp.Dial(fmt.Sprintf("%s:%d", t.FtpSpec.Host, t.FtpSpec.Port))
	res.DelayMs = (time.Now().UnixNano() - start) / 1e6
	if err != nil {
		res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_NET_DIAL
		return &res
	}

	if t.FtpSpec.IfAuth {
		if err := con.Login(t.FtpSpec.Auth.User, t.FtpSpec.Auth.Passwd); err != nil {
			res.Error, res.ErrorCode = err.Error(), pb.TaskResult_ERR_FTP_UNAUTHORIZED
			return &res
		}
	}

	res.Success, res.ErrorCode = true, pb.TaskResult_OK
	return &res
}
