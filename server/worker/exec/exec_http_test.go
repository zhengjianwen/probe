package exec

import (
	"github.com/headzoo/surf"
	"testing"

	"fmt"
	"net/http"
	"time"
)

func TestProbeHttp(t *testing.T) {
	now := time.Now().UnixNano()
	bow := surf.NewBrowser()
	err := bow.Open("http://www.jd.com")
	if err != nil {
		panic(err)
	}

	ms := (time.Now().UnixNano() - now) / 1e6
	fmt.Printf("delay %d\n", ms)
	// Outputs: "The Go Programming Language"
	fmt.Println(bow.Title())
}

func TestProbeHttp_2(t *testing.T) {
	now := time.Now().UnixNano()
	http.Get("127.0.0.1:8080")
	ms := (time.Now().UnixNano() - now) / 1e6
	fmt.Printf("delay %d\n", ms)

}
