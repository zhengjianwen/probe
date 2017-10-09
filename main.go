package main

import (
	"github.com/rongyungo/probe/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
