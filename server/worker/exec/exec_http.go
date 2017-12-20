package exec

import (
	pb "github.com/rongyungo/probe/server/proto"
	"io/ioutil"

	"bytes"
	"io"
	"net"
	"net/http"
	"strings"
	"time"
	"crypto/tls"
)

var trans = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 300 * time.Second,
	}).Dial,
	MaxIdleConns:        50,
	MaxIdleConnsPerHost: 10,
	IdleConnTimeout:     time.Minute * 10,
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

func ProbeHttp(t *pb.Task) *pb.TaskResult {
	start := time.Now().UnixNano()
	err, code, statusCode := DoHttp(t)

	res := ReturnWithCode(t.GetBasicInfo().GetId(), t.GetBasicInfo().GetType(),
		err, start, code)
	res.Http = &pb.TaskResultHttp{
		StatusCode: uint32(statusCode),
	}

	return res
}

func DoHttp(t *pb.Task) (error, pb.TaskResultCode, int) {
	if t.HttpSpec == nil {
		return nil, pb.TaskResult__, 0
	}

	client := &http.Client{Transport: trans}
	req, err := prepareReq(t)
	if err != nil {
		return err, pb.TaskResult_ERR_HTTP_NEW_REQUEST, 0
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err, pb.TaskResult_ERR_HTTP_DO_REQUEST, 0
	}
	defer rsp.Body.Close()

	err, rc := matchRsp(t.HttpSpec.Matcher, rsp)
	return err, rc, rsp.StatusCode
}

func prepareReq(t *pb.Task) (*http.Request, error) {
	spec := t.HttpSpec

	var body io.Reader = nil
	if spec.GetMethod() == pb.HttpSpec_POST && len(spec.GetBody()) > 0 {
		body = bytes.NewBufferString(spec.GetBody())
	}

	req, err := http.NewRequest(spec.Method.String(), spec.Url, body)
	if err != nil {
		return nil, err
	}

	if spec.Header != nil {
		header := make(http.Header)
		for k, v := range spec.Header {
			header.Add(k, v)
		}
		req.Header = header
	}

	if len(spec.Cookies) > 0 {
		req.Header.Set("Cookie", spec.Cookies)
	}

	if spec.BasicAuth != nil {
		req.SetBasicAuth(spec.BasicAuth.User, spec.BasicAuth.Passwd)
	}

	return req, nil
}

func matchRsp(matcher *pb.HttpSpecMatcher, rsp *http.Response) (error, pb.TaskResultCode) {
	if matcher == nil {
		return nil, pb.TaskResult_OK
	}

	if int(matcher.StatusCode) != rsp.StatusCode {
		return ErrStatusCodeUnMatch, pb.TaskResult_ERR_HTTP_STATUS_CODE_UNMATCH
	}

	if len(matcher.Content) > 0 {
		var matched bool
		switch matcher.Target {
		case pb.HttpSpecMatcher_HEAD:
			for key, vl := range rsp.Header {
				if strings.Contains(key, matcher.Content) || strings.Contains(strings.Join(vl, ""), matcher.Content) {
					matched = true
				}
			}

			if matcher.Method == pb.HttpSpecMatcher_EXCLUDE {
				matched = !matched
			}

			if !matched {
				return ErrResponseHeadUnMatch, pb.TaskResult_ERR_HTTP_HEAD_UNMATCH
			}

		case pb.HttpSpecMatcher_BODY:
			data, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				return err, pb.TaskResult_ERR_HTTP_READ_BODY
			}
			body := string(data)
			if strings.ContainsAny(body, matcher.Content) {
				matched = true
			}

			if matcher.Method == pb.HttpSpecMatcher_EXCLUDE {
				matched = !matched
			}

			if !matched {
				return ErrResponseBodyUnMatch, pb.TaskResult_ERR_HTTP_BODY_UNMATCH
			}
		}

	}

	return nil, pb.TaskResult_OK
}
