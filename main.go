package main

import (
	// Force import to test that everything builds correctly
	_ "github.com/pydio/pydio-sdk-go/client"
	"github.com/pydio/pydio-sdk-go/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
