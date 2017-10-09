package exec

import (
	"fmt"
	"testing"
)

func TestProbePing(t *testing.T) {
	if err, n := Pinger("google.com", 8, 8); err != nil {
		t.Fatal(err, n)
	} else {
		fmt.Println(n)
	}
}
