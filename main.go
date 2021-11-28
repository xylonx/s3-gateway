package main

import (
	"os"

	"github.com/xylonx/s3-gateway/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
