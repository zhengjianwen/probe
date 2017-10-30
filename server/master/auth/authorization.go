package auth

import (
	"errors"
)

func CanWriteNode(myid, nodeid int64) (bool, error) {
	var node Node
	err := CallCloud("Srv.NodeInfo", uint64(nodeid), &node)
	if err != nil {
		return false, err
	}

	if node.Type != 1 {
		return false, errors.New("只能在产品线节点配置域名监控")
	}

	var res HasPermissionResponse
	if err := CallUIC("Authority.HasPermission", HasPermissionRequest{
		UserId:    myid,
		NodeId:    nodeid,
		Operation: "tree.modify",
	}, &res); err != nil {
		return false, err
	}

	if res.Has {
		return true, nil
	}
	return false, errors.New("no privileged")
}