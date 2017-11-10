package exec

import (
	//"github.com/headzoo/surf"
	"testing"

	"fmt"
	pb "github.com/rongyungo/probe/server/proto"
)

//
//func TestProbeHttp(t *testing.T) {
//	now := time.Now().UnixNano()
//	bow := surf.NewBrowser()
//	err := bow.Open("http://www.jd.com")
//	if err != nil {
//		panic(err)
//	}
//
//	ms := (time.Now().UnixNano() - now) / 1e6
//	fmt.Printf("delay %d\n", ms)
//	// Outputs: "The Go Programming Language"
//	fmt.Println(bow.Title())
//}

func TestProbeHttp(t *testing.T) {
	task := pb.Task{
		HttpSpec: &pb.HttpSpec{
			Url:    "",
			Method: 1,
		},
	}
	res := ProbeHttp(&task)
	fmt.Printf("%#v\n", res)
}
