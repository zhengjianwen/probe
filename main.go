package main

import (
	"github.com/ten-cloud/prober/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
