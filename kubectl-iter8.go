/*
 */

package main

import (
	"os"

	"github.com/kalantar/iter8r/pkg/cmd"
	"github.com/spf13/pflag"
)

func main() {
	flags := pflag.NewFlagSet("iter8", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.NewCmdIter8Command()
	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
