package main

import (
	"fmt"
	"os"

	cmd "github.com/cwpearson/pipelines-to-docker/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
