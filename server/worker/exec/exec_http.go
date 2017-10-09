package exec

import (
	pb "github.com/ten-cloud/prober/server/proto"
	"io/ioutil"

	"net"
	"net/http"
	"strings"
	"time"
)

var trans = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   5 * time.Second,
		KeepAlive: 20 * time.Second,
	}).Dial,
	MaxIdleConnsPerHost: 1,
}

func ProbeHttp(t *pb.TaskInfo) *pb.TaskResult {
	now := time.Now().UnixNano()
	err, code := DoHttp(t)

	return ReturnWithCode(t.TaskId, err, code, now)
}

func DoHttp(t *pb.TaskInfo) (error, pb.TaskResultCode) {
	if t.Http_Spec == nil {
		return nil, pb.TaskResult__
	}

	client := &http.Client{Transport: trans}
	req, err := prepareReq(t)
	if err != nil {
		return err, pb.TaskResult_ERR_HTTP_NEW_REQUEST
	}

	rsp, err := client.Do(req)
	if err != nil {
		return err, pb.TaskResult_ERR_HTTP_DO_REQUEST
	}

	return matchRsp(t.Http_Spec.Matcher, rsp)
}

func prepareReq(t *pb.TaskInfo) (*http.Request, error) {
	spec := t.Http_Spec

	req, err := http.NewRequest(spec.Method.String(), spec.Url, nil)
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

func matchRsp(matcher *pb.Task_HttpMatcher, rsp *http.Response) (error, pb.TaskResultCode) {
	if matcher == nil {
		return nil, pb.TaskResult_OK
	}

	if int(matcher.StatusCode) != rsp.StatusCode {
		return ErrStatusCodeUnMatch, pb.TaskResult_ERR_HTTP_STATUS_CODE_UNMATCH
	}

	if len(matcher.Content) > 0 {
		var matched bool
		switch matcher.Target {
		case pb.Task_HttpMatcher_HEAD:
			for key, vl := range rsp.Header {
				if strings.Contains(key, matcher.Content) || strings.Contains(strings.Join(vl, ""), matcher.Content) {
					matched = true
				}
			}

			if matcher.Method == pb.Task_HttpMatcher_EXCLUDE {
				matched = !matched
			}

			if !matched {
				return ErrResponseHeadUnMatch, pb.TaskResult_ERR_HTTP_HEAD_UNMATCH
			}

		case pb.Task_HttpMatcher_BODY:
			defer rsp.Body.Close()
			data, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				return err, pb.TaskResult_ERR_HTTP_READ_BODY
			}
			body := string(data)
			if strings.ContainsAny(body, matcher.Content) {
				matched = true
			}

			if matcher.Method == pb.Task_HttpMatcher_EXCLUDE {
				matched = !matched
			}

			if !matched {
				return ErrResponseBodyUnMatch, pb.TaskResult_ERR_HTTP_BODY_UNMATCH
			}
		}

	}

	return nil, pb.TaskResult_OK
}

//func ProberHttpGet(t *pb.TaskInfo) (int32, error) {
//	spec := t.Http_Spec
//
//	httpClient := &http.Client{Transport: trans}
//
//	req, err := http.NewRequest(spec.Method, spec.Url, nil)
//
//	res, err := httpClient.Get(spec.Url)
//	if err != nil {
//		if strings.Contains(err.Error(), "timeout") {
//			return URLErrTimeout, err
//		}
//
//		return URLErrOther, err
//	}
//
//	if res.StatusCode != int(spec.StatusCode) {
//		return URLErrStatusCodeUnmatch, ErrStatusCodeUnMatch
//	}
//
//	if len(spec.BodyMatchText) != 0 {
//		if res.Body == nil {
//			return URLErrResponseBodyUnmatch, ErrStatusCodeUnMatch
//		}
//
//		defer res.Body.Close()
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			return URLErrOther, err
//		}
//
//		content := string(body)
//		if !strings.Contains(content, spec.BodyMatchText) {
//			return URLErrResponseBodyUnmatch, ErrResponseBodyUnMatch
//		}
//	}
//
//	return URLErrNone, nil
//}
