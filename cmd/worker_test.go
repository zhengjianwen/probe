package cmd

import (
	"github.com/satori/go.uuid"
	"testing"

	"fmt"
)

func Test_ss(t *testing.T) {
	a := uuid.NewV3(uuid.UUID("prober"), "1")
	fmt.Println(a)
	//29e36a79-9cd7-11e7-b406-2cd05a808574

}
